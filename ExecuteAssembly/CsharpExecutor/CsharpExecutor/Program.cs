

namespace CsharpExecutor
{
    internal class Program
    {
        public static void Main()
        {
            // 1. 执行本地exe文件，使用Process.start()
            Local.ExeExecutor();
            // 2. 执行本地exe文件(非反射)
            Local.DllExecutor();

        }
        
        
        
    }
}