# YuanTube, a video streaming site demo
As the name suggests, a video streaming site demo, easily adaptable for personal use.
## Architecture Design
![arch](https://github.com/yngyuan/YuanTube/blob/master/arch.png?raw=true)
### Front-end
- Front and back end decoupling
- Login/browse/play page
### Back-end
- Restful API using Go
- Layered processing
- Error handling
- Middleware
### Streaming Service
- Video player module
### File Operation Service
- File operation(Delete, browse) module

## Preview
![video](https://github.com/yngyuan/YuanTube/blob/master/video.png?raw=true)
![index](https://github.com/yngyuan/YuanTube/blob/master/index.png?raw=true)

## TODO
1. ORM abstract layer
2. Safer User Authentication
3. better logging System
