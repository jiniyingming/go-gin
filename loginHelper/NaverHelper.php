<?
namespace app\models\loginHelper;
use app\models\HeadImgDefault;
use Yii;
use yii\helpers\Json;
class NaverHelper extends BaseLogin {

	private $tokenDatas	=	array();

	private $access_token			= '';			// oauth 엑세스 토큰
	private $refresh_token			= '';			// oauth 갱신 토큰
	private $access_token_type		= '';			// oauth 토큰 타입
	private $access_token_expire	= '';			// oauth 토큰 만료
	private $state			= '';			// 네이버 명세에 필요한 검증 키 (현재 버전 라이브러리에서 미검증
	private $loginMode		= 'request';	// 라이브러리 작동 상태
	private $returnCode		= '';			// 네이버에서 리턴 받은 승인 코드
	private $returnState	 = '';			// 네이버에서 리턴 받은 검증 코드
	private $nhnConnectState	= false;
    // action options
	private $autoClose		= false;
	private $curl = NULL;
	private $refreshCount = 1;  // 令牌到期时的续订尝试次数

	public function getLoginUrl(){
        //----请求状态
        $request = Yii::$app->request;
        $nhnMode = $request->get('nhnMode');
        $state = $request->get('state');
        $code= $request->get('code');
        if($nhnMode && $nhnMode != ''){
            $this->loginMode = 'logout';
        }
        if($this->getConnectState() == false){
            $this->generate_state();

            if($state && $code){
                $this->loginMode = 'request_token';
                $this->returnCode = $code;
                $this->returnState = $state;
                $this->_getAccessToken();
            }
        }

		if($this->loginMode == 'request' && (!$this->getConnectState()) ){
            $client_id = Yii::$app->params['NAVER_LOGIN']['client_id'];
            $redirectURI =Yii::$app->params['NAVER_LOGIN']['returnURL'];
            $this->generate_state();
            $state = $this->state;
            $apiURL = "https://nid.naver.com/oauth2.0/authorize?response_type=code&client_id=".$client_id."&redirect_uri=".$redirectURI."&state=".$state;
		    return $apiURL;
		}
        if($this->loginMode == 'request_token'){
            $this->_getAccessToken();
        }
	}
    public function loginCallback()
    {
        $client_id = Yii::$app->params['NAVER_LOGIN']['client_id'];
        $client_secret = Yii::$app->params['NAVER_LOGIN']['client_secret'];
        $redirectURI = Yii::$app->params['NAVER_LOGIN']['returnURL'];
        $get = Yii::$app->request;
        $code = $get->get("code");;
        $state = $get->get("state");;

        $request = Yii::$app->request;
        $nhnMode = $request->get('nhnMode');
        if($nhnMode && $nhnMode != ''){
            $this->loginMode = 'logout';
        }
        if($this->getConnectState() == false){
            $this->generate_state();

            if($state && $code){
                $this->loginMode = 'request_token';
                $this->returnCode = $code;
                $this->returnState = $state;
                $this->_getAccessToken();
            }
        }


        $url = "https://nid.naver.com/oauth2.0/token?grant_type=authorization_code&client_id=" . $client_id . "&client_secret=" . $client_secret . "&redirect_uri=" . $redirectURI . "&code=" . $code . "&state=" . $state;
        $is_post = false;
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_POST, $is_post);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_exec ($ch);
        $status_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        curl_close ($ch);
        $userDate =  $this->getUserProfile();
        if($status_code==200&&is_json($userDate))
        {
            //----返回用户信息
            $userDate = Json::decode($userDate);
            //---请求成功
            if(isset($userDate['result']['message'])&&$userDate['result']['message']==='success'&&isset($userDate['response']))
            {
                $naver_id = $userDate['response']['id'];
//                profile_image  头像索引
                $userData['naver_id'] = $naver_id;
                $userData['info_email'] = isset($userDate['response']['email'])?$userDate['response']['email']:'';
                $userData['username'] =  isset($userDate['response']['name'])?$userDate['response']['name']:'';
                $head_img = HeadImgDefault::find()->select('url')->orderBy('rand()')->asArray()->one();
                $userData['head_img'] = $head_img['url'];
                $this->doCallback(['name'=>'naver_id','value'=>$naver_id,'type'=>'never'],$userData);
            }
        }
    }

	private function getUserProfile($retType = "JSON"){
		if($this->getConnectState()){
			$data = array();
			$data['Authorization'] = $this->access_token_type.' '.$this->access_token;

			$this->curl = curl_init();
			curl_setopt($this->curl, CURLOPT_URL, 'https://apis.naver.com/nidlogin/nid/getUserProfile.xml');
			curl_setopt($this->curl, CURLOPT_POST, 1);
			curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
			curl_setopt($this->curl, CURLOPT_HTTPHEADER, array(
				'Authorization: '.$data['Authorization']
			));

			curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
			$retVar = curl_exec($this->curl);
			curl_close($this->curl);

			$xml = new \SimpleXMLElement($retVar);

			$responseState = (string) $xml->result[0]->resultcode[0];
			if($responseState == "024"){

				if($this->refreshCount > 0){
					$this->refreshCount--;
					$this->_refreshAccessToken();
					$this->getUserProfile();
					return;
				}else{
					return false;
				}
			}

			if($retType == "JSON"){

				$xmlJSON = array();
				$xmlJSON['result']['resultcode'] = (string) $xml->result[0]->resultcode[0];
				$xmlJSON['result']['message'] = (string) $xml->result[0]->message[0];

				if($xml->result[0]->resultcode == '00'){
					foreach($xml->response->children() as $response => $k){
						$xmlJSON['response'][(string)$response] = (string) $k;
					}
				}

				return json_encode($xmlJSON);
			}else{
				return $retVar;
			}
		}else{
			return false;
		}
	}



	/**
	*	Get AccessToken
	*	발급된 엑세스 토큰을 반환합니다. 엑세스 토큰 발급은 로그인 후 자동으로 이루어집니다.
	*/
	function getAccess_token(){
		if($this->access_token){
			return $this->access_token;
		}
	}

	/**
	*	 네이버 연결상태를 반환합니다.
	*    엑세스 토큰 발급/저장이 이루어진 후 connected 상태가 됩니다.
	*/
	function getConnectState(){
		return $this->nhnConnectState;
	}



	private function updateConnectState($strState = ''){
		$this->nhnConnectState = $strState;
	}



    private function _getAccessToken(){
		$data = array();
		$this->curl = curl_init();
		curl_setopt($this->curl, CURLOPT_URL, Yii::$app->params['NAVER_LOGIN']['url'].'token?client_id='.Yii::$app->params['NAVER_LOGIN']['client_id'].'&client_secret='.Yii::$app->params['NAVER_LOGIN']['client_secret'].'&grant_type=authorization_code&code='.$this->returnCode.'&state='.$this->returnState);
		curl_setopt($this->curl, CURLOPT_POST, 1);
		curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
		curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
		$retVar = curl_exec($this->curl);
		curl_close($this->curl);
		$NHNreturns = json_decode($retVar);
		if(isset($NHNreturns->access_token)){

			$this->access_token			= $NHNreturns->access_token;
			$this->access_token_type	= $NHNreturns->token_type;
			$this->refresh_token		= $NHNreturns->refresh_token;
			$this->access_token_expire	= $NHNreturns->expires_in;
			$this->updateConnectState(true);
			if($this->autoClose){
				echo "<script>window.close();</script>";
			}
		}
	}


	private function _refreshAccessToken(){
		$data = array();
		$this->curl = curl_init();
		curl_setopt($this->curl, CURLOPT_URL, Yii::$app->params['NAVER_LOGIN']['url'].'token?client_id='.Yii::$app->params['NAVER_LOGIN']['client_id'].'&client_secret='.Yii::$app->params['NAVER_LOGIN']['client_secret'].'&grant_type=refresh_token&refresh_token='.$this->refresh_token);
		curl_setopt($this->curl, CURLOPT_POST, 1);
		curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
		curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
		$retVar = curl_exec($this->curl);
		curl_close($this->curl);
		$NHNreturns = json_decode($retVar);


		if(isset($NHNreturns->access_token)){


			$this->access_token			= $NHNreturns->access_token;
			$this->access_token_type	= $NHNreturns->token_type;
			$this->access_token_expire	= $NHNreturns->expires_in;

			$this->updateConnectState(true);


		}
	}



  private function generate_state() {
    $mt = microtime();
		$rand = mt_rand();
		$this->state = md5( $mt . $rand );
  }
}
