#include <jni.h>
#include <stdio.h>
#include "NativeExample.h" // 由 javah 工具生成的头文件

JNIEXPORT void JNICALL Java_NativeExample_sayHello(JNIEnv *env, jobject thisObj) {
    printf("Hello from native code!\n");
}