#pragma once  
#include "Point.h"  
class Circle  
{  
private:  
Point center; // 圆心  
double radius; // 半径  

public:  
// 构造函数  
Circle(const Point& center, double radius) : center(center), radius(radius) {}

// 获取圆心  
const Point& getCenter() const { return center; }  

// 获取半径  
double getRadius() const { return radius; }  

// 判断两个圆是否相交  
static bool isIntersecting(const Circle& circle1, const Circle& circle2) {  
	double distance = Point::distance(circle1.getCenter(), circle2.getCenter());  
	return distance <= (circle1.getRadius() + circle2.getRadius());  
}  
};
