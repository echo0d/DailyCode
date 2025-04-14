#pragma once
class Point
{

private:
	double x;
	double y;

public:
	Point(double x, double y) : x(x), y(y)
	{
	}

	static double distance(const Point& point1, const Point& point2)
	{
		return sqrt(pow(point1.x - point2.x, 2) + pow(point1.y - point2.y, 2));
	}
};

