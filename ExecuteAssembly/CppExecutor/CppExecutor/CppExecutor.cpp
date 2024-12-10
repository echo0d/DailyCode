#include <SDKDDKVer.h>

#include <stdio.h>
#include <tchar.h>
#include <windows.h>

#include <metahost.h>
#include <mscoree.h>
#pragma comment(lib, "mscoree.lib")

int _tmain(int argc, _TCHAR* argv[])
{
	ICLRMetaHost* pMetaHost = nullptr;
	ICLRMetaHostPolicy* pMetaHostPolicy = nullptr;
	ICLRRuntimeHost* pRuntimeHost = nullptr;
	ICLRRuntimeInfo* pRuntimeInfo = nullptr;

	HRESULT hr = CLRCreateInstance(CLSID_CLRMetaHost, IID_ICLRMetaHost, (LPVOID*)&pMetaHost);
	hr = pMetaHost->GetRuntime(L"v4.0.30319", IID_PPV_ARGS(&pRuntimeInfo));
	DWORD dwRet = 0;
	if (FAILED(hr))
	{
		goto cleanup;
	}

	hr = pRuntimeInfo->GetInterface(CLSID_CLRRuntimeHost, IID_PPV_ARGS(&pRuntimeHost));

	hr = pRuntimeHost->Start();
	hr = pRuntimeHost->ExecuteInDefaultAppDomain(L"loadCalc.exe",
		L"loadCalc.Program",
		L"bbb",
		L"HELLO!",
		&dwRet);
	hr = pRuntimeHost->Stop();

cleanup:
	if (pRuntimeInfo != nullptr) {
		pRuntimeInfo->Release();
		pRuntimeInfo = nullptr;
	}

	if (pRuntimeHost != nullptr) {
		pRuntimeHost->Release();
		pRuntimeHost = nullptr;
	}

	if (pMetaHost != nullptr) {
		pMetaHost->Release();
		pMetaHost = nullptr;
	}
	return TRUE;
}