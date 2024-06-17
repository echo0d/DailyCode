import type { MenuProps } from 'antd';
import { TabItem } from './AppLayoutFuncs';

//定义每个菜单项的数据对象
export type MenuItem = Required<MenuProps>['items'][number];
/**
 * 根据参数生成菜单项对象
 * @param label 菜单项目显示文字
 * @param key 菜单项唯一key
 * @param icon 菜单项图标
 * @param children 子菜单数组
 * @param type 菜单类型，由子菜单是值为group
 * @returns 返回菜单项实例对象
 */
export function getItem(label: React.ReactNode, key: React.Key, icon?: React.ReactNode, children?: MenuItem[], type?: 'group'): MenuItem {
    //等同于：return {key: key, icon: icon, children: children, label: label, type: type} as MenuItem;
    return {key, icon, children, label, type} as MenuItem;
}
//菜单初始化数据数组
export const menuItems: MenuProps['items'] = [
    getItem('系统菜单', 'systemMenu', null, [
        getItem('功能一', 'systemMenu_menu1'),
        getItem('功能二', 'systemMenu_menu2'),
        getItem('功能三', 'systemMenu_menu3')
    ])
];
//<AppMenu>组件所需父组件传递的数据类型
export type Param = {
    selectedMenuKey: (keyId: string, items: TabItem[], setActiveKey: (newKeyId: string) => void, setItems: (itemArray: TabItem[]) => void) => void,
    setActiveKey: (newKeyId: string) => void,
    items: TabItem[],
    setItems: (itemArray: TabItem[]) => void,
};