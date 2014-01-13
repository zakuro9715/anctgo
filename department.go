package anctgo

type Department string

const (
  Mechanical              Department = "機械工学科"
  Electrical                         = "電気情報工学科"
  Civil                              = "都市システム工学科"
  Architecture                       = "建築学科"
  MechanicalAndElectronic            = "機械・電子システム工学専攻"
  ArchitectureAndCivil               = "建築・都市システム工学専攻"
  M                                  = Mechanical
  E                                  = Electrical
  C                                  = Civil
  A                                  = Architecture
  ME                                 = MechanicalAndElectronic
  AC                                 = ArchitectureAndCivil
)
