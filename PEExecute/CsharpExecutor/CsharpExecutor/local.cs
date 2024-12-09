using System;
using System.Diagnostics;
using System.Reflection;
using System.Runtime.InteropServices;

namespace CsharpExecutor
{
    public class Local
    {
        public static void ExeExecutor()
        {
            // 1. 执行本地exe文件，使用Process.start()
            // 需要执行的文件有参数的话，可以写像下面的processStartInfo，没有参数就执行写文件路径就行。
            string filename = "explorer.exe";
            string arguments = @"C:/";
            ProcessStartInfo processStartInfo = new ProcessStartInfo();
            processStartInfo.FileName = filename;
            processStartInfo.Arguments = arguments;
            Process.Start(processStartInfo);
        }

        // 2. 执行本地dll文件
        public static void DllExecutor()
        {
            [DllImport("user32.dll", CharSet = CharSet.Auto)]
            static extern int MessageBox(IntPtr hWnd, String text, String caption, uint type);

            MessageBox(IntPtr.Zero, "Hello from user32.dll!", "DLL Message", 0);
        }

        // 3.反射加载exe、dll文件(仅限.NET程序集)
        public static void ReflectDllExecutor()
        {
            // 加载 DLL 文件
            Assembly assembly = Assembly.LoadFile(@"C:\Users\echo0d\Desktop\CsharpExecutor\CsharpExecutor\testcalc.dll"); //要绝对路径
            // 获取 DLL 中的类型
            Type type = assembly.GetType("loadCalc.aaa"); //必须使用名称空间+类名称
            
            // 调用类型中的方法
            MethodInfo method = type.GetMethod("bbb"); 
            if (method != null)
            {
                Object obj = assembly.CreateInstance(method.Name);     
                /*// 定义方法参数
                object[] parameters = { "Parameter1", 123 }; // 替换为实际参数值
                // 调用方法并传递参数
                method.Invoke(obj, parameters);*/
                method.Invoke(obj, null);
            }
            else
            {
                Console.WriteLine("Method not found.");
            }
        }
    }
}