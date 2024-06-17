import { TabItem } from "./AppLayoutFuncs";

//标签页点击事件对象
export type TargetKey = React.MouseEvent | React.KeyboardEvent | string;
//<AppTabs>标签需要接收父组件的的数据对象
export type Param = {
    setActiveKey: (newKeyId: string) => void,
    items: TabItem[],
    setItems: (itemArray: TabItem[]) => void,
    activeKey: string
};
/**
 * 点击某页签标题时：将点击的Tab页签状态改为设置状态
 * @param newActiveKey 点击页签的key
 * @param props 父组件传递过来的数据对象
 */
export function onChange (newActiveKey: string, props: Param) {
    props.setActiveKey(newActiveKey);
};
/**
 * 点击某页签的关闭按钮时触发：
 * @param targetKey 被关闭页签的id
 * @param props 父组件传递过来的数据对象
 */
export function remove(targetKey: TargetKey, props: Param) {
    let newActiveKey = props.activeKey;
    let lastIndex = -1;
    props.items.forEach((item, i) => {
        if (item.key === targetKey) {
            lastIndex = i - 1;
        }
    });
    const newPanes = props.items.filter((item) => item.key !== targetKey);
    if (newPanes.length && newActiveKey === targetKey) {
        if (lastIndex >= 0) {
            newActiveKey = newPanes[lastIndex].key;
        } else {
            newActiveKey = newPanes[0].key;
        }
    }
    props.setItems(newPanes);
    props.setActiveKey(newActiveKey);
};
/**
 * 点击某页签的关闭按钮时触发
 * @param targetKey 被关闭的Tab的key
 * @param action 具体动作（add或者remove，由于因此了Tabs的添加按钮，因此onEdit事件只有remove而没有add）
 * @param props 父组件传递过来的数据对象
 */
export function onEdit(targetKey: React.MouseEvent | React.KeyboardEvent | string, action: 'add' | 'remove', props: Param) {
    if (action === 'remove') {
        remove(targetKey, props);
    }
};