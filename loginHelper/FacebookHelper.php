<?php
namespace app\models\loginHelper;
use app\models\ActivityWaitingPoto;
use app\models\Image;
use yii\base\Model;
use Facebook\Facebook;
use Yii;
use Facebook\Exceptions\FacebookResponseException;
use Facebook\Exceptions\FacebookSDKException;
use app\models\User;
use app\models\UserCountry;
use yii\db\Query;
use Zhuzhichao\IpLocationZh\Ip;
use yii\db\Expression;

class FacebookHelper extends BaseLogin
{
    private $config;

    public function init()
    {
        parent::init();
        $this->config = [
            'app_id' => Yii::$app->params['FACEBOOK']['app_id'],
            'app_secret' => Yii::$app->params['FACEBOOK']['app_secret'],
            'persistent_data_handler' => new FacebookPersistentDataHandler()
        ];
    }

    public function getLoginUrl()
    {
        $fb = new Facebook($this->config);
        $helper = $fb->getRedirectLoginHelper();

        if (isset( $_GET['state']))
        {
            $helper->getPersistentDataHandler()->set('state', $_GET['state']);
        }

        $permissions = ['email'];
        $loginUrl = $helper->getLoginUrl(Yii::$app->request->getHostInfo().'/call-back/login-callback?type=fb',$permissions);
        return $loginUrl;
    }

    public function getConnectUrl()
    {
        $fb = new Facebook($this->config);
        $helper = $fb->getRedirectLoginHelper();

        if (isset( $_GET['state']))
        {
            $helper->getPersistentDataHandler()->set('state', $_GET['state']);
        }

        $permissions = ['email'];
        $loginUrl = $helper->getLoginUrl(Yii::$app->request->getHostInfo().'/call-back/login-callback?type=fb',$permissions);
        return $loginUrl;
    }

    public function loginGetPoto($uid)
    {
        $fb = new Facebook($this->config);
        $helper = $fb->getRedirectLoginHelper();
        $permissions = ['email'];
        $loginUrl = $helper->getLoginUrl(Yii::$app->request->getHostInfo().'/call-back/fb-poto?uid='.$uid,$permissions);
        return $loginUrl;
    }

    public function getPoto($uid)
    {
        $fb = new Facebook($this->config);

        $helper = $fb->getRedirectLoginHelper();

        try {
            $accessToken = $helper->getAccessToken();
        } catch(FacebookResponseException $e) {
            // When Graph returns an error
            echo 'Graph returned an error: ' . $e->getMessage();
            exit;
        } catch(FacebookSDKException $e) {
            // When validation fails or other local issues
            echo 'Facebook SDK returned an error: ' . $e->getMessage();
            exit;
        }

        if (! isset($accessToken)) {
            if ($helper->getError()) {
                header('HTTP/1.0 401 Unauthorized');
                echo "Error: " . $helper->getError() . "\n";
                echo "Error Code: " . $helper->getErrorCode() . "\n";
                echo "Error Reason: " . $helper->getErrorReason() . "\n";
                echo "Error Description: " . $helper->getErrorDescription() . "\n";
            } else {
                header('HTTP/1.0 400 Bad Request');
                echo 'Bad request';
            }
            echo "Login failed";
            exit;
        }

        // The OAuth 2.0 client handler helps us manage access tokens
        $oAuth2Client = $fb->getOAuth2Client();

        if (! $accessToken->isLongLived()) {
            try {
                $accessToken = $oAuth2Client->getLongLivedAccessToken($accessToken);
            } catch (FacebookSDKException $e) {
                echo "<p>Error getting long-lived access token: " . $helper->getMessage() . "</p>\n\n";
                exit;
            }
        }

        try {
            // Returns a `Facebook\FacebookResponse` object
            $response = $fb->get('/me?fields=name,picture.type(large)', $accessToken->getValue());
        } catch(FacebookResponseException $e) {
            echo 'Graph returned an error: ' . $e->getMessage();
            exit;
        } catch(FacebookSDKException $e) {
            echo 'Facebook SDK returned an error: ' . $e->getMessage();
            exit;
        }

        $fb_user = $response->getGraphUser();

        $new_waiting = new ActivityWaitingPoto();

        $new_waiting->uid = $uid;

        $new_waiting->name = $fb_user['name'];

        $new_waiting->poto = $fb_user['picture']['url'];

        $new_waiting->create_time = time();

        $new_waiting->save();
    }

    public function connectCallback($uid)
    {
        $fb = new Facebook($this->config);

        $helper = $fb->getRedirectLoginHelper();

        try {
            $accessToken = $helper->getAccessToken();
        } catch(FacebookResponseException $e) {
            // When Graph returns an error
            echo 'Graph returned an error: ' . $e->getMessage();
            exit;
        } catch(FacebookSDKException $e) {
            // When validation fails or other local issues
            echo 'Facebook SDK returned an error: ' . $e->getMessage();
            exit;
        }

        if (! isset($accessToken)) {
            if ($helper->getError()) {
                header('HTTP/1.0 401 Unauthorized');
                echo "Error: " . $helper->getError() . "\n";
                echo "Error Code: " . $helper->getErrorCode() . "\n";
                echo "Error Reason: " . $helper->getErrorReason() . "\n";
                echo "Error Description: " . $helper->getErrorDescription() . "\n";
            } else {
                header('HTTP/1.0 400 Bad Request');
                echo 'Bad request';
            }
            echo "<script>self.close();</script>";
            exit;
        }

        // The OAuth 2.0 client handler helps us manage access tokens
        $oAuth2Client = $fb->getOAuth2Client();

        if (! $accessToken->isLongLived()) {
            try {
                $accessToken = $oAuth2Client->getLongLivedAccessToken($accessToken);
            } catch (FacebookSDKException $e) {
                echo "<p>Error getting long-lived access token: " . $helper->getMessage() . "</p>\n\n";
                exit;
            }
        }

        try {
            // Returns a `Facebook\FacebookResponse` object
            $response = $fb->get('/me?fields=id,name,age_range,gender,locale,picture.type(large),email', $accessToken->getValue());
        } catch(FacebookResponseException $e) {
            echo 'Graph returned an error: ' . $e->getMessage();
            exit;
        } catch(FacebookSDKException $e) {
            echo 'Facebook SDK returned an error: ' . $e->getMessage();
            exit;
        }

        $fb_user = $response->getGraphUser();

        $facebook_id = isset($fb_user['id'])?trim($fb_user['id']):'';

        if(empty($facebook_id)) die('Not available FaceBook_ID');
        $res = (new Query())->from('{{%user}}')->where(['facebook_id'=>$facebook_id])->one();
        if(!empty($res)){
            //请绑定为使用的账号
            echo '<script>self.close();window.opener.alert('.lang("sorry, this account is already used, please use another").');</script>';
        }else{
            $user = (new Query())->from('{{%user}}')->select(['facebook_id'])->where(['id'=>$uid])->one();
            if(!empty($user['google_id'])){
                echo '<script>self.close();window.opener.alert("sorry, this account is already connected");</script>';
            }else{
                Yii::$app->db->createCommand()->update('{{%user}}',['facebook_id'=>$facebook_id],['id'=>$uid])->execute();
                echo '<script>self.close();window.opener.location.reload();</script>';
            }

        }
    }

    public function loginCallback()
    {
        $fb = new Facebook($this->config);

        $helper = $fb->getRedirectLoginHelper();

        try {
            $accessToken = $helper->getAccessToken(Yii::$app->request->getHostInfo().'/call-back/login-callback?type=fb');
        } catch(FacebookResponseException $e) {
            // When Graph returns an error
            echo 'Graph returned an error: ' . $e->getMessage();
            exit;
        } catch(FacebookSDKException $e) {
            // When validation fails or other local issues
            echo 'Facebook SDK returned an error: ' . $e->getMessage();
            exit;
        }

        if (! isset($accessToken)) {
            if ($helper->getError()) {
                header('HTTP/1.0 401 Unauthorized');
                echo "Error: " . $helper->getError() . "\n";
                echo "Error Code: " . $helper->getErrorCode() . "\n";
                echo "Error Reason: " . $helper->getErrorReason() . "\n";
                echo "Error Description: " . $helper->getErrorDescription() . "\n";
            } else {
                header('HTTP/1.0 400 Bad Request');
                echo 'Bad request';
            }
            echo "<script>self.close();</script>";
            exit;
        }

        // The OAuth 2.0 client handler helps us manage access tokens
        $oAuth2Client = $fb->getOAuth2Client();

        if (! $accessToken->isLongLived()) {
            try {
                $accessToken = $oAuth2Client->getLongLivedAccessToken($accessToken);
            } catch (FacebookSDKException $e) {
                echo "<p>Error getting long-lived access token: " . $helper->getMessage() . "</p>\n\n";
                exit;
            }
        }

        try {
            // Returns a `Facebook\FacebookResponse` object
            $response = $fb->get('/me?fields=id,name,age_range,gender,locale,picture.type(large),email', $accessToken->getValue());
        } catch(FacebookResponseException $e) {
            echo 'Graph returned an error: ' . $e->getMessage();
            exit;
        } catch(FacebookSDKException $e) {
            echo 'Facebook SDK returned an error: ' . $e->getMessage();
            exit;
        }

        $fb_user = $response->getGraphUser();

        $facebook_id = isset($fb_user['id'])?trim($fb_user['id']):'';

        if(empty($facebook_id)) die('Not available FaceBook_ID');
        $userInfo['username'] = (isset($fb_user['name']) && !empty($fb_user['name']))?trim($fb_user['name']):'default';
//        $userInfo['head_img'] = (isset($fb_user['picture']['url']) && !empty($fb_user['picture']['url']))?trim($fb_user['picture']['url']):'';
        $userInfo['info_email'] = (isset($fb_user['email']) && !empty($fb_user['email']))?trim($fb_user['email']):'';
        $userInfo['locale'] = (isset($fb_user['locale']) && !empty($fb_user['locale']))?strtolower(str_replace('-','_',trim($fb_user['locale']))):'';
        $this->doCallback(['name'=>'facebook_id','value'=>$facebook_id,'type'=>'facebook'],$userInfo);
    }
}
