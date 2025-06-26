/*
#define OLC_PGE_APPLICATION
#include "PixelGameEngine.h"
#include "PixelGameEngine.h"
// https://github.com/OneLoneCoder
// https://github.com/OneLoneCoder/olcPixelGameEngine/wiki
// https://www.youtube.com/watch?v=DDV_2cWT94U&list=PLrOv9FMX8xJEEQFU_eAvYTsyudrewymNl
// Override base class with your custom functionality
class Example : public olc::PixelGameEngine
{
public:
	Example()
	{
		// Name your application
		sAppName = "Example";
	}

public:
	bool OnUserCreate() override
	{
		// Called once at the start, so create things here
		return true;
	}

	bool OnUserUpdate(float fElapsedTime) override
	{
		// Called once per frame, draws random coloured pixels
		for (int x = 0; x < ScreenWidth(); x++)
			for (int y = 0; y < ScreenHeight(); y++)
				Draw(x, y, olc::Pixel(rand() % 256, rand() % 256, rand() % 256));
		return true;
	}
};

int main()
{
	Example demo;
	if (demo.Construct(256, 240, 4, 4))
		demo.Start();
	return 0;
}


#include <iostream>
#include <string>
#include <queue>
using namespace std;


#define N 4
#define M 5

struct Node {

	int val, i, index;
};

struct comp {

	bool operator() (const Node &lhr, const Node &lhs) const {
		return lhr.val > lhs.val;
	}
};

void showSort(int list[N][]) {

	priority_queue<Node, std::vector<Node>, comp> q;

	for (int i = 0; i < M; i++)
		q.push({list[i][0], i, 0});
}


template<class T>
void showSort(T data[], int n) {
	for(int i = 0; i < n; i++){
		std::cout << data[i]  << " ";
	}
}

template<class T>
void merge(T data[], T temp[], int low, int middle, int high) {
	for(int i = low; i <= high; i++) {
		temp[i] = data[i];
	}

	int i = low;
	int j = middle + 1;
	int k = low;
}
