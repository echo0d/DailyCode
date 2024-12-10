# 把代码写进ps1脚本里

$Assemblies = (
    "System, Version=2.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089, processorArchitecture=MSIL",
    "System.Linq, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a, processorArchitecture=MSIL"
)

$Source = @"
using System;
using System.Reflection;
namespace TestApplication
{
    public class Program
    {
        public static void Main()
        {

            Console.WriteLine("HELLO");
        }
    }
}
"@

Add-Type -ReferencedAssemblies $Assemblies -TypeDefinition $Source -Language CSharp
[TestApplication.Program]::Main()