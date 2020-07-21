<?
namespace app\models\loginHelper;
use Yii;

class Naver extends BaseLogin {

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
	private $autoClose		= true;
	private $curl = NULL;
	private $refreshCount = 1;  // 토큰 만료시 갱신시도 횟수
	public function getLoginUrl(){
		if($this->loginMode == 'request' && (!$this->getConnectState()) ){
            $client_id = Yii::$app->params['NAVER_LOGIN']['client_id'];
            $redirectURI =Yii::$app->params['NAVER_LOGIN']['returnURL'];
            $this->generate_state();
            $state = $this->state;
            $apiURL = "https://nid.naver.com/oauth2.0/authorize?response_type=code&client_id=".$client_id."&redirect_uri=".$redirectURI."&state=".$state;
		    return $apiURL;
            $api = Yii::$app->params['NAVER_LOGIN']['url'].'authorize?client_id='.Yii::$app->params['NAVER_LOGIN']['client_id'].'&response_type=code&redirect_uri='.Yii::$app->params['NAVER_LOGIN']['returnURL'].'&state='.$this->state;
            return $api;
		}
	}

	function logout(){
		$this->refreshCount = 1;
		$data = array();
		$this->curl = curl_init();
		curl_setopt($this->curl, CURLOPT_URL, Yii::$app->params['NAVER_LOGIN']['url'].'token?client_id='.Yii::$app->params['NAVER_LOGIN']['client_id'].'&client_secret='.Yii::$app->params['NAVER_LOGIN']['client_secret'].'&grant_type=delete&refresh_token='.$this->refresh_token.'&sercive_provider=NAVER');
		curl_setopt($this->curl, CURLOPT_POST, 1);
		curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
		curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
		$retVar = curl_exec($this->curl);
		curl_close($this->curl);
		echo "<script>window.location.href = 'http://".$_SERVER["HTTP_HOST"] . $_SERVER['PHP_SELF']."';</script>";
	}


	function getUserProfile($retType = "JSON"){
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
					$this->logout();
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
