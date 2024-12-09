package org.example;

import java.io.IOException;

public class Main {
    public static void main(String[] args) {
        try {
            // 创建进程
            ProcessBuilder processBuilder = new ProcessBuilder(".\\loadCalc.exe");
            processBuilder.redirectErrorStream(true); // 合并错误流
            Process process = processBuilder.start();

            // 等待进程结束
            int exitCode = process.waitFor();
            System.out.println("Exited with code: " + exitCode);
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}