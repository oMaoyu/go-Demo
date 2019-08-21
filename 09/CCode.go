package _9

/*
#include <windows.h>	// 引入头文件
#include <conio.h>
#include <stdtdb.h>

// 使用Windows API移动控制台的光标位置
void gotoxy(int x,int y)
{
    COORD c;	// <windows.h> 结构体，表示一个字符在控制台屏幕上的坐标。
    c.X=x, c.Y=y;
    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);	// 修改位置
}

// 清空控制台回显
int direct()
{
    return _getch();
}

// 去除控制台位置提示光标 “_”
void hideCursor()
{
	CONSOLE_CURSOR_INFO  cci;		// 定义光标类对象
	cci.bVisible = FALSE;			// 设光标属性为 "不可见"
	cci.dwSize = sizeof(cci);
	SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE), &cci);	// 重设光标信息。
}
*/
import "C" 			// Go中可以直接使用C语言函数

//重设控制台光标位置
func GotoPosition(X int, Y int) {
	C.gotoxy(C.int(X), C.int(Y))	//调用C函数。语法：C.函数名
}

//获取键盘输入字符，隐藏回显。
func Direction() (key int) {
	key = int(C.direct())
	return
}

//隐藏控制台光标
func HideCursor() {
	C.hideCursor()
}
