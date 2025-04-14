// GeometryUtils.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include "Circle.h"

int main() {
    // 提示用户输入第一个圆的圆心坐标和半径
    double x1, y1, r1;
    std::cout << "Enter the center (x, y) and radius of the first circle: ";
    std::cin >> x1 >> y1 >> r1;

    // 提示用户输入第二个圆的圆心坐标和半径
    double x2, y2, r2;
    std::cout << "Enter the center (x, y) and radius of the second circle: ";
    std::cin >> x2 >> y2 >> r2;

    // 创建两个圆形对象
    Circle circle1(Point(x1, y1), r1);
    Circle circle2(Point(x2, y2), r2);

    // 判断两个圆是否相交
    if (Circle::isIntersecting(circle1, circle2)) {
        std::cout << "The two circles intersect." << std::endl;
    }
    else {
        std::cout << "The two circles do not intersect." << std::endl;
    }

    return 0;
}