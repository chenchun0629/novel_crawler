# novel_crawler

hello_world.proto <-> api(dto) <- service(dto->do) <- biz(do) (依赖倒置)-> data(do->po)

# 层级
- api: 为 grpc/http 的服务定义服务的service/client interface， dto对象都会使用 protobuf工具生产
- service: 实现api的service interface 处理dto->do对象的转换，
- biz: 业务逻辑层，定义do对象需要使用repo，建议do使用贫血模型，
- data: 存储层

# 对象

- dto: data transfer object 数据传输对象
- do: domain object 领域对象，负责业务领域实体逻辑的行为对象载体
- po: persistent object 持久化对象

