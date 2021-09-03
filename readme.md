# Printlog

- 简介：用来做终端及文件日志输出，轻量简便多端指定
- 功能：
  - ● 日志级别分为5级，输出，警告，debug，错误，异常五种
  - ● 输出方式分为终端和日志文件，终端带有彩色打印，和秒级时间戳输出
  - ● 区分开发者模式和生产模式，生产模式可以关掉 ”debug“ 将不再输出
- 属性：
  - IsDebug 是否开启Debug若不开启Debug将不会输出
  - OutFile 指定日志输出路径及名称，默认 "为项目根路径/月日时分秒.log"
- 使用方法：
  1. 导入报名`import github.com/ggcodec/Printlog`（如果觉得包名长可以起别名）
  2. 使用NewPringlog() 得到一个log实例，通过log实例操作
  3. 方法分为两类一类Fmt方法向终端输出日志，第二类File方法向文件输入

注意：请谨慎使用FmtPanic 和 FilePanic 这会让你的程序造成恐慌并退出
