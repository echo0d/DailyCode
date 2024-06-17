import { CheckboxChangeEvent } from 'antd/es/checkbox';

//任务列表项数据结构
export type TaskListItem = {
    taskName: string,
    checkedFlag: boolean
};
//state中的业务数据的数据结构
export type DataInfoType = {
    taskName: string,
    taskList: TaskListItem[]
};

/**
 * 输入框onChange事件处理函数
 * @param newTaskname 修改后的值
 * @param dataInfo state中的业务数据
 * @param setDataInfo 赋值修改值
 */
export function taskNameChange(newTaskname: string, dataInfo: DataInfoType, setDataInfo: (value: DataInfoType) => void) {
    let newStateInfo = {...dataInfo};
    newStateInfo.taskName = newTaskname;
    setDataInfo(newStateInfo);
}

/**
 * 添加按钮点击事件处理函数
 * @param dataInfo state中的业务数据
 * @param setDataInfo state修改值函数
 */
export  function addNewTaskToList(dataInfo: DataInfoType, setDataInfo: (value: DataInfoType) => void) {
    let newStateInfo = { ...dataInfo };
    newStateInfo.taskList.push({
      taskName: newStateInfo.taskName,
      checkedFlag: false,
    });
    // 添加任务后，清空输入框中的输入内容
    newStateInfo.taskName = "";
    setDataInfo(newStateInfo);
  }

/**
 * Checkbox onChange事件绑定函数
 * @param e CheckboxOnChange事件对象
 * @param dataInfo 
 * @param setDataInfo 
 */
export function checkedOrNotchecked(e: CheckboxChangeEvent, dataInfo: DataInfoType, setDataInfo: (value: DataInfoType) => void) {
  let newStateInfo = { ...dataInfo };
  newStateInfo.taskList.map((item) => {
    if (item.taskName === e.target.value) {
      return (item.checkedFlag = e.target.checked);
    } else {
      return false;
    }
  });
  setDataInfo(newStateInfo);
}

/**
 * 删除按钮处理函数
 * @param dataInfo state数据对象
 * @param setDataInfo state值修改函数
 */
export function deleteCheckedTaskName(dataInfo: DataInfoType, setDataInfo: (value: DataInfoType) => void) {
  let newStateInfo = { ...dataInfo };
  newStateInfo.taskList = newStateInfo.taskList.filter(
    (item) => !item.checkedFlag
  );
  setDataInfo(newStateInfo);
};