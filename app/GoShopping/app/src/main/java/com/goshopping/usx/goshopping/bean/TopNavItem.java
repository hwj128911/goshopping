package com.goshopping.usx.goshopping.bean;

/**
 * Created by Si on 2016/10/19.
 */

   /*顶部导航实体类*/
public class TopNavItem {

    private int name;
    private int iconRes;

    public TopNavItem(int name, int iconRes) {
        this.name = name;
        this.iconRes = iconRes;
    }

    public int getName() {
        return name;
    }

    public void setName(int name) {
        this.name = name;
    }

    public int getIconRes() {
        return iconRes;
    }

    public void setIconRes(int iconRes) {
        this.iconRes = iconRes;
    }
}
