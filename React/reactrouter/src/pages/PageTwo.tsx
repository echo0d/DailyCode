import React from 'react';
import { useLoaderData } from 'react-router-dom';
//新增数据加载函数，函数根据url传递过来的参数，加载返回数据
export async function loader(urlBizData: any) {
    console.log(urlBizData.params.bizDataName);
    return {name: '查收数据库的数据'};//返回的示例数据，实际项目请根据场景进行修改
  }

function PageTwo() {
    //获取的就是加载完成的数据
    const bizDataInfo = useLoaderData();
    console.log(bizDataInfo); //打印的就是loader返回的数据
    return (
        <React.Fragment>
            <h1 style={{textAlign: 'center', color: 'red'}}>页面组件二</h1>
        </React.Fragment>
    );
}

export default PageTwo;