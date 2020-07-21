<?php
namespace app\models\loginHelper;
use yii\base\Model;
use Abraham\TwitterOAuth\TwitterOAuth;
use Yii;
use app\models\User;
use app\models\UserCountry;
use yii\db\Query;
use yii\web\Cookie;
use Zhuzhichao\IpLocationZh\Ip;
use yii\db\Query as db;
use yii\db\Expression;

class TwitterHelper extends BaseLogin
{

    public function getLoginUrl()
    {
        $tt = new TwitterOAuth(Yii::$app->params['TWITTER']['CONSUMER_KEY'], Yii::$app->params['TWITTER']['CONSUMER_SECRET']);

        $request_token = $tt->oauth('oauth/request_token', [
            'oauth_callback' => 'https://pngtree.com/call-back/login-callback?type=tt'
        ]);

        $cookie = Yii::$app->response->cookies;
        $cookie->add(new Cookie(['name'=>'tt','value'=>$request_token['oauth_token'],'expire'=>time()+300,'domain'=>".pngtree.com"]));
        $cookie->add(new Cookie(['name'=>'tts','value'=>$request_token['oauth_token_secret'],'expire'=>time()+300,'domain'=>".pngtree.com"]));

        $loginUrl = $tt->url('oauth/authorize', array('oauth_token' => $request_token['oauth_token']));

        return $loginUrl;
    }

    public function getConnectUrl()
    {
        $tt = new TwitterOAuth(Yii::$app->params['TWITTER']['CONSUMER_KEY'], Yii::$app->params['TWITTER']['CONSUMER_SECRET']);

        $request_token = $tt->oauth('oauth/request_token', [
            'oauth_callback' => 'https://pngtree.com/call-back/login-callback?type=tt'
        ]);

        $cookie = Yii::$app->response->cookies;
        $cookie->add(new Cookie(['name'=>'tt','value'=>$request_token['oauth_token'],'expire'=>time()+300,'domain'=>".pngtree.com"]));
        $cookie->add(new Cookie(['name'=>'tts','value'=>$request_token['oauth_token_secret'],'expire'=>time()+300,'domain'=>".pngtree.com"]));

        $loginUrl = $tt->url('oauth/authorize', array('oauth_token' => $request_token['oauth_token']));

        return $loginUrl;
    }

    /**
     * twitter回调
     */
    public function connectCallback($uid)
    {
        $oauth_token = Yii::$app->request->get('oauth_token');
        if(!$oauth_token){
            if(Yii::$app->request->get('denied') !== null){
                echo "<script>self.close();</script>";
                exit;
            }else{
                die('Illegal request.');
            }
        }

        $request_token = [];
        $cookie = Yii::$app->request->cookies;
//        $request_token['oauth_token'] = Yii::$app->session->get('tt_oauth_token');
//        $request_token['oauth_token_secret'] = Yii::$app->session->get('tt_oauth_token_secret');
        $request_token['oauth_token'] = $cookie->getValue('tt');
        $request_token['oauth_token_secret'] = $cookie->getValue('tts');
        $cookie2 = Yii::$app->response->cookies;
        $cookie2->remove('tt');
        $cookie2->remove('tts');

        if ($oauth_token != $request_token['oauth_token']) die('Illegal request.');

        $tt = new TwitterOAuth(Yii::$app->params['TWITTER']['CONSUMER_KEY'], Yii::$app->params['TWITTER']['CONSUMER_SECRET'], $request_token['oauth_token'], $request_token['oauth_token_secret']);

        $access_token = $tt->oauth("oauth/access_token", ["oauth_verifier" => Yii::$app->request->get('oauth_verifier')]);

        $tt = new TwitterOAuth(Yii::$app->params['TWITTER']['CONSUMER_KEY'], Yii::$app->params['TWITTER']['CONSUMER_SECRET'], $access_token['oauth_token'], $access_token['oauth_token_secret']);

        $params = ['include_email' => 'true', 'include_entities' => 'false', 'skip_status' => 'true'];

        $tt_user = $tt->get("account/verify_credentials",$params);

        $twitter_id = isset($tt_user->id_str)?trim($tt_user->id_str):'';

        if(empty($twitter_id)) die('Not available twitter_id');

        $res = (new Query())->from('{{%user}}')->where(['twitter_id'=>$twitter_id])->one();
        if(!empty($res)){
            //请绑定为使用的账号
            echo '<script>self.close();window.opener.alert('.lang("sorry, this account is already used, please use another").');</script>';
        }else{
            $user = (new Query())->from('{{%user}}')->select(['twitter_id'])->where(['id'=>$uid])->one();
            if(!empty($user['twitter_id'])){
                echo '<script>self.close();window.opener.alert("sorry, this account is already connected");</script>';
            }else{
                Yii::$app->db->createCommand()->update('{{%user}}',['twitter_id'=>$twitter_id],['id'=>$uid])->execute();
                echo '<script>self.close();window.opener.location.reload();</script>';
            }

        }

    }

    /**
     * twitter回调
     */
    public function loginCallback()
    {
        $oauth_token = Yii::$app->request->get('oauth_token');
        if(!$oauth_token){
            if(Yii::$app->request->get('denied') !== null){
                echo "<script>self.close();</script>";
                exit;
            }else{
                die('Illegal request.');
            }
        }

        $request_token = [];
        $cookie = Yii::$app->request->cookies;
//        $request_token['oauth_token'] = Yii::$app->session->get('tt_oauth_token');
//        $request_token['oauth_token_secret'] = Yii::$app->session->get('tt_oauth_token_secret');
        $request_token['oauth_token'] = $cookie->getValue('tt');
        $request_token['oauth_token_secret'] = $cookie->getValue('tts');
        $cookie2 = Yii::$app->response->cookies;
        $cookie2->remove('tt');
        $cookie2->remove('tts');

        if ($oauth_token != $request_token['oauth_token']) die('Illegal request.');

        $tt = new TwitterOAuth(Yii::$app->params['TWITTER']['CONSUMER_KEY'], Yii::$app->params['TWITTER']['CONSUMER_SECRET'], $request_token['oauth_token'], $request_token['oauth_token_secret']);

        $access_token = $tt->oauth("oauth/access_token", ["oauth_verifier" => Yii::$app->request->get('oauth_verifier')]);

        $tt = new TwitterOAuth(Yii::$app->params['TWITTER']['CONSUMER_KEY'], Yii::$app->params['TWITTER']['CONSUMER_SECRET'], $access_token['oauth_token'], $access_token['oauth_token_secret']);

        $params = ['include_email' => 'true', 'include_entities' => 'false', 'skip_status' => 'true'];

        $tt_user = $tt->get("account/verify_credentials",$params);

        $twitter_id = isset($tt_user->id_str)?trim($tt_user->id_str):'';

        if(empty($twitter_id)) die('Not available twitter_id');

        $userInfo['username'] = (isset($tt_user->name) && !empty($tt_user->name))?trim($tt_user->name):'default';
//        $userInfo['head_img'] = (isset($tt_user->profile_image_url_https) && !empty($tt_user->profile_image_url_https))?str_replace('_normal','',trim($tt_user->profile_image_url_https)):'';
        $userInfo['info_email'] = (isset($tt_user->email) && $tt_user->email !== null)?trim($tt_user->email):'';
        $userInfo['locale'] = (isset($tt_user->lang) && !empty($tt_user->lang))?strtolower(str_replace('-','_',trim($tt_user->lang))):'';
        $this->doCallback(['name'=>'twitter_id','value'=>$twitter_id,'type'=>'twitter'],$userInfo);
    }

    /**
     * 注册登陆处理
     */
    public function doYinqinrenRecall($twitter,$userInfo){
        $this->doCallback($twitter,$userInfo);
    }
}
