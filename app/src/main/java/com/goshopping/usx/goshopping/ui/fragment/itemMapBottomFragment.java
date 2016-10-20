package com.goshopping.usx.goshopping.ui.fragment;

import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import com.goshopping.usx.goshopping.R;

import java.util.Map;

/**
 * Created by hwj on 2016-10-17.
 */

public class itemMapBottomFragment extends Fragment{


    private View mConvertView;
    private Map<String,String> mViewMessage;
    private TextView titleMessageTv;
    private TextView summaryMessageTv;

    public void setViewMessage(Map<String,String> viewMessage){
        mViewMessage = viewMessage;
    }

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {

        mConvertView = inflater.inflate(R.layout.item_bottom_map,container,false);
        initView();
        initData();
        return mConvertView;
    }

    private void initView() {
        titleMessageTv = (TextView) mConvertView.findViewById(R.id.title_item_map_bottom);
        summaryMessageTv = (TextView) mConvertView.findViewById(R.id.summary_item_map_bottom);
    }

    private void initData() {
        if(mViewMessage == null) return;

        String title = mViewMessage.get("title");
        String summary = mViewMessage.get("summary");
        if( title != null){
            titleMessageTv.setText(title);
        }
        if(summary != null){
            summaryMessageTv.setText(summary);
        }
    }

}
