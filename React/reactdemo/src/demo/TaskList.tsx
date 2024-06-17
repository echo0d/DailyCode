import React from 'react';
import { Checkbox } from 'antd';
import { CheckboxChangeEvent } from 'antd/es/checkbox';
//定义接收父组件传递过来的state中的业务数据类型
type ParamsType = {
    taskDataInfo: {
        taskName: string,
        taskList: {
            taskName: string,
            checkedFlag: boolean
        }[]
    }
    //增加定义接收函数类型的参数，
    checkboxChange: (event: CheckboxChangeEvent) => void
};

//函数式组件中的参数，父组件传递过来的业务数据
function TaskList(props: ParamsType) {
    return (
        <React.Fragment>
            {
                // 遍历父组件state中的taskList数据，生成<li>标签
                props.taskDataInfo.taskList.map(oneItem => {
                    return (
                        <li key={oneItem.taskName}>
                            <Checkbox checked={oneItem.checkedFlag} value={oneItem.taskName} onChange={props.checkboxChange}>
                                {oneItem.taskName}
                            </Checkbox>
                            
                        </li>
                    )
                })
            }
        </React.Fragment>
    );
}

export default TaskList;