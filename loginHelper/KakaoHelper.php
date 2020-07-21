<?php
namespace app\models\loginHelper;
use Abraham\TwitterOAuth\TwitterOAuth;
use app\models\HeadImgDefault;
use Yii;

use yii\web\Cookie;

class KakaoHelper extends BaseLogin{
    private $tokenDatas	=	array();
    private $access_token			= '';
    private $refresh_token			= '';
    private $access_token_type		= '';
    private $access_token_expire	= '';
    private $client_id		= '';
    private $client_secret	= '';
    private $returnURL		= '';
    private $state			= '';

    private $encode_state = 'n';
    private $loginMode		= 'request';
    private $returnCode		= '';
    private $returnState	 = '';
    private $kakaoConnectState	= false;
    // action options
    private $autoClose		= true;
    private $showLogout		= true;
    private $curl = NULL;
    private $refreshCount = 1;  //



    public  function  getLoginUrl()
    {
        $redirect_uri = Yii::$app->request->getHostInfo().'/call-back/login-callback?type=kakao';
        $Secretkey = Yii::$app->params['KAKAO']['api'];
        $Secreturl = Yii::$app->params['KAKAO']['url'];
        $kakaourl = sprintf($Secreturl.'oauth/authorize?client_id=%s'.'&redirect_uri=%s'.'&response_type=%s',$Secretkey,$redirect_uri,'code');
        return $kakaourl;
    }

    public function loginCallback(){

        $this->generate_state();

        if($_GET['code']){
            $this->loginMode = 'request_token';
            $this->returnCode = $_GET['code'];
            $this->returnState = Yii::$app->request->get('state',0);
        }

        $data = array();
        $this->curl = curl_init();
        curl_setopt($this->curl, CURLOPT_URL, Yii::$app->params['KAKAO']['url'].'oauth/token?client_id='.Yii::$app->params['KAKAO']['api'].'&grant_type=authorization_code&code='.$this->returnCode);
        curl_setopt($this->curl, CURLOPT_POST, 1);
        curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
        curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
        $retVar = curl_exec($this->curl);
        curl_close($this->curl);
        $KAKAOreturns = json_decode($retVar);
        if(isset($KAKAOreturns->access_token)){
            $this->access_token		= $KAKAOreturns->access_token;
            $this->access_token_type	= $KAKAOreturns->token_type;
            $this->refresh_token		= $KAKAOreturns->refresh_token;
            $this->access_token_expire	= $KAKAOreturns->expires_in;
            $this->updateConnectState(true);

        }
        $kakaoData = $this->getUserProfile();
        $this->doCallback(['name'=>'kakao_id','value'=>$kakaoData['kakao_id'],'type'=>'kakao'],$kakaoData);
    }
    function getConnectState(){
        return $this->kakaoConnectState;
    }
    function getUserProfile($retType = "JSON"){
        if($this->getConnectState()){
            $data = array();
            $userData = [];
            $data['Authorization'] = $this->access_token_type.' '.$this->access_token;
            $this->curl = curl_init();
            curl_setopt($this->curl, CURLOPT_URL, 'https://kapi.kakao.com/v1/user/me');
            curl_setopt($this->curl, CURLOPT_POST, 1);
            curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
            curl_setopt($this->curl, CURLOPT_HTTPHEADER, array(
                'Authorization: '.$data['Authorization'],
            ));
            curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
            $retVar = curl_exec($this->curl);
            curl_close($this->curl);

            $_retAr = json_decode($retVar);
            $_retAr = json_decode(json_encode($_retAr),true);

            if(!$_retAr['id']){

                if($this->refreshCount > 0){
                    $this->refreshCount--;
                    $this->_refreshAccessToken();
                    $this->getUserProfile();
                    return;
                }else{
                    return false;
                }
            }
            if (!empty($_retAr))
            {
                if(!isset($_retAr['id']))return false;
                $userData['kakao_id'] = $_retAr['id'];
                $userData['info_email'] = isset($_retAr['kaccount_email'])?$_retAr['kaccount_email']:'';
                $userData['username'] = $_retAr['properties']['nickname'];
//                if(isset($_retAr['properties']['profile_image'])&&!empty($_retAr['properties']['profile_image']))
//                {
//                    $userData['head_img'] = $_retAr['properties']['profile_image'];
//                }else{
                $head_img = HeadImgDefault::find()->select('url')->orderBy('rand()')->asArray()->one();
                $userData['head_img'] = $head_img['url'];
//                }

            }

            return $userData;
        }else{
            return false;
        }
    }
    private function _refreshAccessToken(){
        $data = array();
        $this->curl = curl_init();
        curl_setopt($this->curl, CURLOPT_URL, Yii::$app->params['KAKAO']['url'].'oauth/token?client_id='.Yii::$app->params['KAKAO']['api'].'&grant_type=refresh_token&refresh_token='.$this->refresh_token);
        curl_setopt($this->curl, CURLOPT_POST, 1);
        curl_setopt($this->curl, CURLOPT_POSTFIELDS, $data);
        curl_setopt($this->curl, CURLOPT_RETURNTRANSFER,true);
        $retVar = curl_exec($this->curl);
        curl_close($this->curl);
        $KAKAOreturns = json_decode($retVar);
        if(isset($KAKAOreturns->access_token)){
            $this->access_token			= $KAKAOreturns->access_token;
            $this->access_token_type	= $KAKAOreturns->token_type;
            $this->access_token_expire	= $KAKAOreturns->expires_in;
            $this->updateConnectState(true);
        }
    }

    private function updateConnectState($strState = ''){
        $this->kakaoConnectState = $strState;
    }
    private function generate_state() {
        $mt = microtime();
        $rand = mt_rand();
        $this->state = md5( $mt . $rand );
    }

    function arrayXML($data, &$xml_ip){

        foreach( $data as $key => $value ) {

            if( is_array($value) ) {

                if( is_numeric($key) ){
                    $key = 'item'.$key;
                }

                $subnode = $xml_ip->addChild($key);
                $this->arrayXML($value, $subnode);
            } else {
                $xml_ip->addChild("$key",htmlspecialchars("$value"));
            }
        }
    }
}
