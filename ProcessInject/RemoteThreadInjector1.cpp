#include <windows.h>
#include <iostream>

// DWORD: 32位无符号整型，typedef unsigned long DWORD;
// HANDLE: Windows对象句柄，typedef void* HANDLE;
// LPVOID: 通用指针类型，typedef void* LPVOID;
// LPTHREAD_START_ROUTINE: 线程函数指针类型，typedef DWORD (WINAPI *LPTHREAD_START_ROUTINE)(LPVOID);

int main(int argc, char *argv[])
{
    SetConsoleOutputCP(CP_UTF8);

    if (argc != 3)
    {
        std::cout << "用法: " << argv[0] << " <进程ID> <DLL路径>" << std::endl;
        return 1;
    }

    DWORD pid = std::stoul(argv[1]);
    const char *dllPath = argv[2];

    // 打开目标进程
    HANDLE hProcess = OpenProcess(PROCESS_ALL_ACCESS, FALSE, pid);
    if (!hProcess)
    {
        std::cerr << "无法打开进程: " << GetLastError() << std::endl;
        return 1;
    }

    // 在目标进程中分配内存
    LPVOID pRemoteBuf = VirtualAllocEx(hProcess, NULL, strlen(dllPath) + 1,
                                       MEM_COMMIT, PAGE_READWRITE);
    if (!pRemoteBuf)
    {
        std::cerr << "VirtualAllocEx 失败: " << GetLastError() << std::endl;
        CloseHandle(hProcess);
        return 1;
    }

    // 写入 DLL 路径到目标进程
    if (!WriteProcessMemory(hProcess, pRemoteBuf, dllPath, strlen(dllPath) + 1, NULL))
    {
        std::cerr << "WriteProcessMemory 失败: " << GetLastError() << std::endl;
        VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
        CloseHandle(hProcess);
        return 1;
    }

    // 获取 LoadLibraryA 地址
    LPTHREAD_START_ROUTINE pfnThreadRtn = (LPTHREAD_START_ROUTINE)
        GetProcAddress(GetModuleHandleA("Kernel32"), "LoadLibraryA");
    if (!pfnThreadRtn)
    {
        std::cerr << "GetProcAddress 失败: " << GetLastError() << std::endl;
        VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
        CloseHandle(hProcess);
        return 1;
    }

    // 创建远程线程
    HANDLE hThread = CreateRemoteThread(hProcess, NULL, 0, pfnThreadRtn,
                                        pRemoteBuf, 0, NULL);
    if (!hThread)
    {
        std::cerr << "CreateRemoteThread 失败: " << GetLastError() << std::endl;
        VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
        CloseHandle(hProcess);
        return 1;
    }

    // 等待线程结束
    WaitForSingleObject(hThread, INFINITE);

    // 清理
    VirtualFreeEx(hProcess, pRemoteBuf, 0, MEM_RELEASE);
    CloseHandle(hThread);
    CloseHandle(hProcess);

    std::cout << "DLL 注入成功！" << std::endl;
    return 0;
}