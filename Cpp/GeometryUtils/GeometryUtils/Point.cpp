#include "Point.h"
#include <iostream>

Point::~Point()
{
}
Point::Point(double x, double y) : x(x), y(y)
{
}
// 计算两点之间的欧几里得距离
double Point::distance(const Point& point1, const Point& point2)
{
	return sqrt(pow(point1.x - point2.x, 2) + pow(point1.y - point2.y, 2));
}