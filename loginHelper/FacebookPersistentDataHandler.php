<?php
/**
 * Created by PhpStorm.
 * User: simbapeng
 * Date: 17/3/5
 * Time: 上午12:47
 */

namespace app\models\loginHelper;
use Facebook\PersistentData\PersistentDataInterface;
use Yii;
use yii\web\Cookie;


class FacebookPersistentDataHandler implements PersistentDataInterface
{
    /**
     * @var string Prefix to use for session variables.
     */
    protected $sessionPrefix = 'FBRLH_';

    public function get($key)
    {
        return Yii::$app->request->cookies->getValue($this->sessionPrefix .$key);
    }

    public function set($key, $value)
    {

        if($value===null){
            Yii::$app->response->cookies->remove($this->sessionPrefix .$key);
        }else{
            Yii::$app->response->cookies->add(new Cookie(['name'=>($this->sessionPrefix .$key),'value'=>$value]));
        }
    }
}
