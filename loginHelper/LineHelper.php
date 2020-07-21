<?php
namespace app\models\loginHelper;
use app\models\HeadImgDefault;
use Yii;

use yii\web\Cookie;

class LineHelper extends BaseLogin{

    /**
     * 獲取登陸連接
     */
    public  function  getLoginUrl()
    {
        $state = sha1(time());
        Yii::$app->response->cookies->add(new Cookie([
            'name' => 'l_line',
            'value' => $state,
            'expire'=>time()+90,
            'domain'=>".pngtree.com"
        ]));
        $base_url = "https://access.line.me/oauth2/v2.1/authorize";
        $client_id = Yii::$app->params['LINE']['client_id'];
        $redirect_uri = Yii::$app->request->getHostInfo().'/call-back/line';
        $query = "";
        $query .= "response_type=" . urlencode("code") . "&";
        $query .= "client_id=" . urlencode($client_id) . "&";
        $query .= "redirect_uri=" . urlencode($redirect_uri) . "&";
        $query .= "state=" . urlencode($state) . "&";
//        $query .= "scope=" . "&";
        $query .= "scope=profile%20openid%20email"  . "&";
//        $query .= "nonce=1234"  . "&";
        $url = $base_url . '?' . $query;
        return $url;
    }

    public function loginCallback($code){
        $obj = json_decode($this->getToken($code), true);
        if(!isset($obj['id_token']))
        {
            echo '<script>self.close();alert("sorry, this account is already connected");window.location.href='/';</script>';
            return;
        }
        $val = explode(".", $obj['id_token']);
        $data_json = base64_decode($val[1]);
        $data = json_decode($data_json,true);
        if(!isset($data['sub'])||empty($data['sub'])) die('Not available LINE_ID,please try again');
        $userInfo = [];
        $userInfo['username'] = (isset($data['name']) && !empty($data['name']))?trim($data['name']):'default';
        $head_img = HeadImgDefault::find()->select('url')->orderBy('rand()')->asArray()->one();
        $userInfo['head_img'] = $head_img['url'];
        $userInfo['info_email'] = (isset($data['email']) && !empty($data['email']))?trim($data['email']):'';
        $this->doCallback(['name'=>'line_id','value'=>$data['sub'],'type'=>'line'],$userInfo);
    }

    public function getToken($code){
        $client_id = Yii::$app->params['LINE']['client_id'];
        $client_secret = Yii::$app->params['LINE']['client_secret'];
        $redirect_uri = Yii::$app->request->getHostInfo()."/call-back/line";

        $curl = curl_init();
        curl_setopt_array($curl, array(
            CURLOPT_URL => "https://api.line.me/oauth2/v2.1/token",
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
            CURLOPT_CUSTOMREQUEST => "POST",
            CURLOPT_POSTFIELDS => "grant_type=authorization_code&code=" . $code . "&client_id=" . $client_id . "&client_secret=" . $client_secret . "&redirect_uri=" . $redirect_uri,
            CURLOPT_HTTPHEADER => array(
                "cache-control: no-cache",
                "content-type: application/x-www-form-urlencoded"
            ),
        ));
        $response = curl_exec($curl);
        curl_close($curl);
        return $response;
    }

}
