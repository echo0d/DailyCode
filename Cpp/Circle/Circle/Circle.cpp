#include <iostream>

class circle
{
private:
    /* data */
    float radius; //圆的半径
    float high; //圆的高
    float area; //圆的面积
    float circumference; //圆的周长
    float volume; //圆的体积
    float surface; //圆的表面积
    static const float PI; // 静态常量 圆周率

public:
    circle(float r = 0, float h = 0); //构造函数
    ~circle();

    // 计算圆周长
    float calculateCircumference();

    // 计算圆面积
    float calculateArea();

    // 计算圆球表面积
    float calculateSphereArea();

    // 计算圆柱体积
    float calculateCylinderVolume();
};

// 在类外初始化静态常量
const float circle::PI = 3.14;

circle::circle(float r, float h)
{
    radius = r;
    high = h;
}

circle::~circle()
{
}

float circle::calculateCircumference()
{
    circumference = 2 * PI * radius;
    return circumference;
}

float circle::calculateArea()
{
    area = PI * radius * radius;
    return area;
}

float circle::calculateSphereArea()
{
    surface = 4 * PI * radius * radius;
    return surface;
}

float circle::calculateCylinderVolume()
{
    volume = PI * radius * radius * high;
    return volume;
}

int main()
{
    float radius, high;

    // 输入半径和高
    std::cin >> radius >> high;

    // 创建circle对象并设置值
    circle c(radius, high);

    // 计算并输出结果，以空格分隔
    std::cout << c.calculateCircumference() << " "
        << c.calculateArea() << " "
        << c.calculateSphereArea() << " "
        << c.calculateCylinderVolume() << std::endl;

    return 0;
}
