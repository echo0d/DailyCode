// import React from 'react';
//import { AppstoreOutlined, MailOutlined, SettingOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';//引入菜单项的数据类型
import { Menu } from 'antd';
import {Param, menuItems} from './AppMenuFuncs';

// type MenuItem = Required<MenuProps>['items'][number];
// function getItem(label: React.ReactNode, key: React.Key, icon?: React.ReactNode, children?: MenuItem[], type?: 'group'): MenuItem {
//     return {key, icon, children, label, type} as MenuItem;
// }
// const items: MenuProps['items'] = [
//     getItem('系统菜单', 'systemMenu', null, [
//         getItem('功能一', 'systemMenu_menu1'),
//         getItem('功能二', 'systemMenu_menu2'),
//         getItem('功能三', 'systemMenu_menu3')
//     ])
// ];

function AppMenu(props: Param) {
    const onClick: MenuProps['onSelect'] = (e) => {
        // console.log('click ', e);
        props.selectedMenuKey(e.key, props.items, props.setActiveKey, props.setItems);
    };
    return (
        <Menu onSelect={onClick} defaultSelectedKeys={['systemMenu_menu1']} defaultOpenKeys={['systemMenu']} mode="inline" items={menuItems} theme='dark' />
    );
}

export default AppMenu;