//定义Tabs的每个标签页对象的数据结构
export type TabItem = { label: string, children: string, key: string, closable?: boolean };
//定义Tabs的数组，并提供首页页签的默认数据
export const initialTabItems = [
    { label: '首页', children: '首页页签内容', key: '1', closable: false }
];
/**
 * 选中某菜单项时调用函数，在Tabs中添加新的Tab
 * @param menuKeyId 选中的菜单数据keyId
 * @param items state中的Tab对象数组
 * @param setActiveKey 修改Tabs选中状态的Tabid的函数
 * @param setItems 修改state中的Tab对象数组的函数
 */
export function selectedMenuKey(menuKeyId: string, items: TabItem[], setActiveKey: (newKeyId: string) => void, setItems: (itemArray: TabItem[]) => void) {
    addNewTab(menuKeyId, '功能：' + menuKeyId, items, setActiveKey, setItems);
}
/**
 * 添加新的Tab：如果点击菜单对应的Tab页签已经存在，则将Tab页签设置为选中状态，否则添加新的页签并设置为选中状态
 * @param tabKey 新Tab对象的key
 * @param tabLabel 新Tab对象的Title文字
 * @param items state中的Tabs数组
 * @param setActiveKey 修改Tabs选中状态的Tabid的函数
 * @param setItems 修改state中的Tab对象数组的函数
 */
export function addNewTab(tabKey: string, tabLabel: string, items: TabItem[], setActiveKey: (newKeyId: string) => void, setItems: (itemArray: TabItem[]) => void) {
    if (items.some((oneItem) => oneItem.key === tabKey)) {
        setActiveKey(tabKey);
    } else {
        const newPanes = [...items];
        newPanes.push({ label: tabLabel, children: tabKey + "，Tab内容就是：" + tabLabel, key: tabKey });
        setItems(newPanes);
        setActiveKey(tabKey);
    }
};