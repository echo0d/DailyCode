using System;

namespace loadCalc
{
    public class Program
    {
        public static void Main()
        {
            Console.WriteLine("Hello World!");
            System.Diagnostics.Process p = new System.Diagnostics.Process();
            p.StartInfo.FileName = "c:\\windows\\system32\\calc.exe";
            p.Start();
        }
    }
    public class aaa
    {
        public static void bbb()
        {
            System.Diagnostics.Process p = new System.Diagnostics.Process();
            p.StartInfo.FileName = "c:\\windows\\system32\\calc.exe";
            p.Start();
        }
    }
}