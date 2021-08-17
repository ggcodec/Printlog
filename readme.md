# Printlog,输出日志工具



- 简介：用来做终端及文件日志输出，轻量简便多端指定
- 功能：Linux下带有彩色终端输出日志，带有秒级时间戳，
- 属性：
  - IsDebug 是否开启Debug若不开启Debug将不会输出
  - OutFile 指定日志输出路径及名称，默认 "为项目根路径/月日时分秒.log"
- 使用方法：
  1. 导入报名`import lo github.com/ggcodec/Printlog`（如果觉得包名长可以起别名）
  2. 使用NewPringlog() 得到一个log实例，通过log实例操作
  3. 方法分为两类一类Fmt方法向终端输出日志，第二类File方法向文件输入
