
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

# 模型
- 失血模型: 仅仅包含数据的定义和getter/setter，java中的pojo，.net中的poco****
- 贫血模型: 包含一些业务逻辑，但不包含持久层的逻辑。这部分依赖于持久层的逻辑将会放到服务层中。贫血模型中的领域对象是不依赖持久层的。
- 充血模型: 包含所有的业务逻辑，包含依赖持久层的业务逻辑。
- 胀血模型: 把业务逻辑不想管的其他逻辑都放在领域模型中。另一种的嗜血模型，因为服务层消失了，领域层干了服务层的事情，到头来什么都没改变。

