package com.goshopping.usx.goshopping;

import android.app.Application;

/**
 * Created by Si on 2016/10/14.
 */

public class GoShoppingApp extends Application {

    public static GoShoppingApp mInstance;


    @Override
    public void onCreate()
    {

        super.onCreate();

        mInstance = this;

    }

    public static GoShoppingApp getInstance()
    {

        return mInstance;
    }
}
