#pragma once  
#include "Point.h"  
class Circle  
{  
private:  
Point center; // Բ��  
double radius; // �뾶  

public:  
// ���캯��  
Circle(const Point& center, double radius) : center(center), radius(radius) {}

// ��ȡԲ��  
const Point& getCenter() const { return center; }  

// ��ȡ�뾶  
double getRadius() const { return radius; }  

// �ж�����Բ�Ƿ��ཻ  
static bool isIntersecting(const Circle& circle1, const Circle& circle2) {  
	double distance = Point::distance(circle1.getCenter(), circle2.getCenter());  
	return distance <= (circle1.getRadius() + circle2.getRadius());  
}  
};
