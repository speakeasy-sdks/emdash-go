# JumpGate


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ConnectedSystems`                                                   | [][ConnectedSystem](../../models/shared/connectedsystem.md)          | :heavy_check_mark:                                                   | The systems within range of the gate that have a corresponding gate. |
| `FactionSymbol`                                                      | **string*                                                            | :heavy_minus_sign:                                                   | The symbol of the faction that owns the gate.                        |
| `JumpRange`                                                          | *float64*                                                            | :heavy_check_mark:                                                   | The maximum jump range of the gate.                                  |