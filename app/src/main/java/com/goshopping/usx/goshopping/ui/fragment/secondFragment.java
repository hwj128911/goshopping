package com.goshopping.usx.goshopping.ui.fragment;

import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.goshopping.usx.goshopping.R;

/**
 * Created by hwj on 2016-10-15.
 */

public class secondFragment extends Fragment {

    private View mConvertView;

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {

        if(mConvertView == null){
            mConvertView = inflater.inflate(R.layout.second_fragment_content,container,false);

            initView();
        }
        ViewGroup parent = (ViewGroup) mConvertView.getParent();
        if (parent != null){
            parent.removeView(mConvertView);
        }
        return mConvertView;

    }

    private void initView() {
    }
}
