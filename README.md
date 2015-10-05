# srtAdjuster
Movie subscript (.srt type) time axis adjuster - add or subtract time by golang.

## Editor
Atom 1.0.19

## Design
### Modules
分为两个文件-主文件和工具文件；
Two files:  helper, main
#### helper
工具文件提供两个方法：将特定的字符串转换为千分秒的数字；
另一个方法是将千分秒的数字转换为字符串，如：

#### main
主函数读取命令行参数：偏移的毫秒数-为整数；正数向后偏，即将对话向后推移；
负数向前偏，即将对话向前提。

然后读取字幕文件，一行行处理，读一行，写一行：写至另一个文件中；
如果是符合要求的行，如：00:00:00,000 --> 00:00:10,000
就开始处理-做偏移操作。


## test
./testHelper.sh

## Usage
./run.sh as an example

当然，go 是跨平台的语言；可以在Windows上使用。如编译成windows上的可执行文件srtAdjuster，然后：

srtAdjuster.exe -filename ./testData/tintin-partial-01.txt -offset 2000
