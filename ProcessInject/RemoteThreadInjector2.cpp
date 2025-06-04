#include <windows.h>
#include <tlhelp32.h>
#include <iostream>

// DWORD: 32位无符号整型，typedef unsigned long DWORD;
// HANDLE: Windows对象句柄，typedef void* HANDLE;
// LPVOID: 通用指针类型，typedef void* LPVOID;
// LPTHREAD_START_ROUTINE: 线程函数指针类型，typedef DWORD (WINAPI *LPTHREAD_START_ROUTINE)(LPVOID);

// 获取指定进程名的第一个进程ID
DWORD GetProcessIdByName(const char *processName)
{
    DWORD processId = 0;
    PROCESSENTRY32 pe32;
    pe32.dwSize = sizeof(PROCESSENTRY32);
    HANDLE hProcessSnap = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
    if (hProcessSnap == INVALID_HANDLE_VALUE)
        return 0;
    if (Process32First(hProcessSnap, &pe32))
    {
        do
        {
            if (strcmp(pe32.szExeFile, processName) == 0)
            {
                processId = pe32.th32ProcessID;
                break;
            }
        } while (Process32Next(hProcessSnap, &pe32));
    }
    CloseHandle(hProcessSnap);
    return processId;
}

int main(int argc, char *argv[])
{
    SetConsoleOutputCP(CP_UTF8);
    // 检查参数数量
    if (argc != 3)
    {
        std::cout << "用法: " << argv[0] << " <进程名.exe> <DLL路径>" << std::endl;
        return 1;
    }

    // 获取目标进程ID
    DWORD pid = GetProcessIdByName(argv[1]);
    if (pid == 0)
    {
        std::cerr << "未找到指定进程: " << argv[1] << std::endl;
        return 1;
    }

    // 打开目标进程
    HANDLE hProcess = OpenProcess(PROCESS_ALL_ACCESS, FALSE, pid);
    if (!hProcess)
    {
        std::cerr << "无法打开进程: " << GetLastError() << std::endl;
        return 1;
    }

    // 获取DLL路径
    const char *dllPath = argv[2];
    // 在目标进程中分配内存
    LPVOID pRemoteBuf = VirtualAllocEx(hProcess, NULL, strlen(dllPath) + 1, MEM_COMMIT, PAGE_READWRITE);
    if (!pRemoteBuf)
    {
        std::cerr << "VirtualAllocEx 失败: " << GetLastError() << std::endl;
        CloseHandle(hProcess);
        return 1;
    }
    // 写入DLL路径到目标进程
    if (!WriteProcessMemory(hProcess, pRemoteBuf, dllPath, strlen(dllPath) + 1, NULL))
    {
        std::cerr << "WriteProcessMemory 失败: " << GetLastError() << std::endl;
        VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
        CloseHandle(hProcess);
        return 1;
    }
    // 获取LoadLibraryA地址
    LPTHREAD_START_ROUTINE pfnThreadRtn = (LPTHREAD_START_ROUTINE)GetProcAddress(GetModuleHandleA("Kernel32"), "LoadLibraryA");
    if (!pfnThreadRtn)
    {
        std::cerr << "GetProcAddress 失败: " << GetLastError() << std::endl;
        VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
        CloseHandle(hProcess);
        return 1;
    }
    // 创建远程线程，加载DLL
    HANDLE hThread = CreateRemoteThread(hProcess, NULL, 0, pfnThreadRtn, pRemoteBuf, 0, NULL);
    if (!hThread)
    {
        std::cerr << "CreateRemoteThread 失败: " << GetLastError() << std::endl;
        VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
        CloseHandle(hProcess);
        return 1;
    }
    // 等待远程线程执行完毕
    WaitForSingleObject(hThread, INFINITE);
    // 释放分配的内存和句柄
    VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
    CloseHandle(hThread);
    CloseHandle(hProcess);
    std::cout << "DLL 注入成功！" << std::endl;
    return 0;
}