<?php
namespace app\models\loginHelper;
use Google_Client;
use Google_Service_People;
use Google_Service_Oauth2;
use Yii;
use yii\db\Query;

class GoogleHelper extends BaseLogin
{
    public function getLoginUrl()
    {
        $client = new Google_Client();
        $client->setAuthConfig(Yii::$app->params['GOOGLE']['config_file']);
        $client->setAccessType("offline");        // offline access
        $client->setIncludeGrantedScopes(true);   // incremental auth
        $client->addScope([
            Google_Service_People::USERINFO_PROFILE,
            Google_Service_People::USERINFO_EMAIL
        ]);
        $client->setRedirectUri(Yii::$app->request->getHostInfo() . '/call-back/login-callback?type=gg');

        $auth_url = $client->createAuthUrl();
        $loginUrl = filter_var($auth_url, FILTER_SANITIZE_URL);

        return $loginUrl;
    }

    public function getConnectUrl()
    {
        $client = new Google_Client();
        $client->setAuthConfig(Yii::$app->params['GOOGLE']['config_file']);
        $client->setAccessType("offline");        // offline access
        $client->setIncludeGrantedScopes(true);   // incremental auth
        $client->addScope([
            Google_Service_People::USERINFO_PROFILE,
            Google_Service_People::USERINFO_EMAIL
        ]);
        $client->setRedirectUri(Yii::$app->request->getHostInfo() . '/call-back/login-callback?type=gg');

        $auth_url = $client->createAuthUrl();
        $loginUrl = filter_var($auth_url, FILTER_SANITIZE_URL);

        return $loginUrl;
    }

    /**
     * twitter回调
     */
    public function connectCallback($uid)
    {
        $outh_code = Yii::$app->request->get('code');
        if(!$outh_code) die('Illegal request.');

        $client = new Google_Client();
        $client->setAuthConfig(Yii::$app->params['GOOGLE']['config_file']);
        $client->authenticate($outh_code);
        $access_token = $client->getAccessToken();
        if(!$access_token)
        {
            echo '<script>self.close();alert("sorry, this account is already connected");window.location.href='/';</script>';
            return;
        }
        $client->setAccessToken($access_token);

        $server = new Google_Service_Oauth2($client);

        $gg_user = $server->userinfo->get();

        $google_id = isset($gg_user->id)?trim($gg_user->id):'';

        if(empty($google_id)) die('Not available google_id');

        $res = (new Query())->from('{{%user}}')->where(['google_id'=>$google_id])->one();
        if(!empty($res)){
            //请绑定为使用的账号
            echo '<script>self.close();window.opener.alert('.lang("sorry, this account is already used, please use another").');</script>';
        }else{
            $user = (new Query())->from('{{%user}}')->select(['google_id'])->where(['id'=>$uid])->one();
            if(!empty($user['google_id'])){
                echo '<script>self.close();window.opener.alert("sorry, this account is already connected");</script>';
            }
            Yii::$app->db->createCommand()->update('{{%user}}',['google_id'=>$google_id],['id'=>$uid])->execute();
            echo '<script>self.close();window.opener.location.reload();</script>';
        }

    }

    /**
     * google回调处理
     */
    public function loginCallback()
    {
        $outh_code = Yii::$app->request->get('code');
        if(!$outh_code) die('Illegal request.');

        $client = new Google_Client();
        $client->setAuthConfig(Yii::$app->params['GOOGLE']['config_file']);
        $client->authenticate($outh_code);
        $access_token = $client->getAccessToken();

        $client->setAccessToken($access_token);

        $server = new Google_Service_Oauth2($client);

        $gg_user = $server->userinfo->get();

        $google_id = isset($gg_user->id)?trim($gg_user->id):'';

        if(empty($google_id)) die('Not available google_id');
        $userInfo['username'] = (isset($gg_user->name) && !empty($gg_user->name))?trim($gg_user->name):'default';
        $userInfo['info_email'] = (isset($gg_user->email) && !empty($gg_user->email))?trim($gg_user->email):'';
        $userInfo['locale'] = (isset($gg_user->locale) && !empty($gg_user->locale))?strtolower(str_replace('-','_',trim($gg_user->locale))):'';
        $this->doCallback(['name'=>'google_id','value'=>$google_id,'type'=>'google'],$userInfo);
    }
}
