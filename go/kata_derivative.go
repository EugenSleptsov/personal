package kata

import "regexp"
import "fmt"
import "strconv"
import "math"

type Type string
const (
	Operator Type = "Operator"
	Function Type = "Function"
	Constant Type = "Constant"
	Variable Type = "Variable"
)

type Expression struct {
	Type Type
	Value interface{}
	ArgumentOne *Expression
	ArgumentTwo *Expression
}

func (expr *Expression) Reduce() *Expression {
	switch expr.Type {
	case Function:
		expr.ArgumentOne = expr.ArgumentOne.Reduce()
	case Operator:
		expr.ArgumentOne = expr.ArgumentOne.Reduce()
		expr.ArgumentTwo = expr.ArgumentTwo.Reduce()

		switch expr.Value.(string) {
		case "+":
			if expr.ArgumentOne.Type == Constant {
        if expr.ArgumentOne.Value.(float64) == 0.0 {
          expr = expr.ArgumentTwo
        } else if expr.ArgumentTwo.Type == Constant {
					expr = &Expression{
						Type: Constant,
						Value: expr.ArgumentOne.Value.(float64) + expr.ArgumentTwo.Value.(float64),
					}
				} else if expr.ArgumentOne.Type != Operator { // operators cannot be summed
					if expr.ArgumentOne.Type == expr.ArgumentTwo.Type {
						expr = &Expression{
							Type: Operator,
							Value: "*",
							ArgumentOne: &Expression{
								Type: Constant,
								Value: 2.0,
							},
							ArgumentTwo: expr.ArgumentOne,
						}
					} else if expr.ArgumentTwo.Type == Operator && expr.ArgumentTwo.Value.(string) == "*" && expr.ArgumentTwo.ArgumentOne.Type == Constant && expr.ArgumentTwo.ArgumentTwo.Type == expr.ArgumentOne.Type {
						expr = &Expression{
							Type: Operator,
							Value: "*",
							ArgumentOne: &Expression{
								Type: Constant,
								Value: expr.ArgumentTwo.ArgumentOne.Value.(float64) + 1.0,
							},
							ArgumentTwo: expr.ArgumentTwo.ArgumentTwo,
						}
					}
				}
			}
		case "-":
			if expr.ArgumentOne.Type == Constant {
				if expr.ArgumentTwo.Type == Constant {
					expr = &Expression{
						Type: Constant,
						Value: expr.ArgumentOne.Value.(float64) - expr.ArgumentTwo.Value.(float64),
					}
				}
			}
		case "*":
			if expr.ArgumentOne.Type == Constant && expr.ArgumentOne.Value.(float64) == 1.0 {
				expr = expr.ArgumentTwo
			} else if expr.ArgumentTwo.Type == Constant && expr.ArgumentTwo.Value.(float64) == 1.0 {
				expr = expr.ArgumentOne
			} else if expr.ArgumentTwo.Type == Constant && expr.ArgumentTwo.Value.(float64) == 0.0 || expr.ArgumentOne.Type == Constant && expr.ArgumentOne.Value.(float64) == 0.0 {
				expr = &Expression{
					Type: Constant,
					Value: 0.0,
				}
			} else if expr.ArgumentOne.Type == Constant && expr.ArgumentTwo.Type == Operator && expr.ArgumentTwo.Value.(string) == "*" && expr.ArgumentTwo.ArgumentOne.Type == Constant {
        expr = &Expression{
          Type: Operator,
          Value: "*",
          ArgumentOne: &Expression{
            Type: Constant,
            Value: expr.ArgumentOne.Value.(float64) * expr.ArgumentTwo.ArgumentOne.Value.(float64),
          },
          ArgumentTwo: expr.ArgumentTwo.ArgumentTwo,
        }
      }
		case "/":
      if expr.ArgumentOne.Type == Constant && expr.ArgumentTwo.Type == Constant {
        expr = &Expression{
          Type: Constant,
          Value: expr.ArgumentOne.Value.(float64) / expr.ArgumentTwo.Value.(float64),
        }
      }
		case "^":
			if expr.ArgumentTwo.Type == Constant {
				if expr.ArgumentTwo.Value.(float64) == 1.0 {
					expr = expr.ArgumentOne
				} else if expr.ArgumentOne.Type == Constant {
					expr = &Expression{
						Type: Constant,
						Value: math.Pow(expr.ArgumentOne.Value.(float64), expr.ArgumentTwo.Value.(float64)),
					}
				}
			}
		}
	}

	return expr
}

func (expr *Expression) Diff() *Expression {
	if expr.Type == Constant {
		return &Expression{
			Type: Constant,
			Value: 0.0,
		}
	} else if expr.Type == Variable {
		return &Expression{
			Type: Constant,
			Value: 1.0,
		}
	} else if expr.Type == Operator {
		switch expr.Value.(string) {
		case "+","-":
			return &Expression{
				Type: expr.Type,
				Value: expr.Value.(string),
				ArgumentOne: expr.ArgumentOne.Diff(),
				ArgumentTwo: expr.ArgumentTwo.Diff(),
			}
		case "/":
			return &Expression{
				Type: Operator,
				Value: "/",
				ArgumentOne: &Expression{
					Type: Operator,
					Value: "-",
					ArgumentOne: &Expression{
						Type: Operator,
						Value: "*",
						ArgumentOne: expr.ArgumentOne.Diff(),
						ArgumentTwo: expr.ArgumentTwo,
					},
					ArgumentTwo: &Expression{
						Type: Operator,
						Value: "*",
						ArgumentOne: expr.ArgumentOne,
						ArgumentTwo: expr.ArgumentTwo.Diff(),
					},
				},
				ArgumentTwo: &Expression{
					Type: Operator,
					Value: "^",
					ArgumentOne: expr.ArgumentTwo,
					ArgumentTwo: &Expression{
						Type: Constant,
						Value: 2.0,
					},
				},
			}
		case "*":
			return &Expression{
				Type: Operator,
				Value: "+",
				ArgumentOne: &Expression{
					Type: Operator,
					Value: "*",
					ArgumentOne: expr.ArgumentOne.Diff(),
					ArgumentTwo: expr.ArgumentTwo,
				},
				ArgumentTwo: &Expression{
					Type: Operator,
					Value: "*",
					ArgumentOne: expr.ArgumentOne,
					ArgumentTwo: expr.ArgumentTwo.Diff(),
				},
			}
		case "^":
			if expr.ArgumentTwo.Type == Constant {
				return &Expression{
					Type: Operator,
					Value: "*",
					ArgumentOne: &Expression{
						Type: Constant,
						Value: expr.ArgumentTwo.Value.(float64),
					},
					ArgumentTwo: &Expression{
						Type: Operator,
						Value: "*",
						ArgumentOne: expr.ArgumentOne.Diff(),
						ArgumentTwo: &Expression{
							Type: Operator,
							Value: "^",
							ArgumentOne: expr.ArgumentOne,
							ArgumentTwo: &Expression{
								Type: Constant,
								Value: expr.ArgumentTwo.Value.(float64) - 1.0,
							},
						},
					},
				}
			} else {
				return nil
			}
		}
	} else if expr.Type == Function {
		// chain rule
		var ExprMain, ExprInside *Expression

		ExprInside = expr.ArgumentOne.Diff()
		switch expr.Value {
		case "cos":
			ExprMain = &Expression{
				Type: Operator,
				Value: "*",
				ArgumentOne: &Expression{
					Type: Constant,
					Value: -1.0,
				},
				ArgumentTwo: &Expression{
					Type: Function,
					Value: "sin",
					ArgumentOne: expr.ArgumentOne,
				},
			}
		case "sin":
			ExprMain = &Expression{
				Type: Function,
				Value: "cos",
				ArgumentOne: expr.ArgumentOne,
			}
		case "tan":
			ExprMain = &Expression{
				Type: Operator,
				Value: "+",
				ArgumentOne: &Expression{
					Type: Constant,
					Value: 1.0,
				},
				ArgumentTwo: &Expression{
					Type: Operator,
					Value: "^",
					ArgumentOne: &Expression{
						Type: Function,
						Value: "tan",
						ArgumentOne: expr.ArgumentOne,
					},
					ArgumentTwo: &Expression{
						Type: Constant,
						Value: 2.0,
					},
				},
			}
		case "exp":
			ExprMain = &Expression{
				Type: Function,
				Value: "exp",
				ArgumentOne: expr.ArgumentOne,
			}
		case "ln":
			ExprMain = &Expression{
				Type: Operator,
				Value: "/",
				ArgumentOne: &Expression{
					Type: Constant,
					Value: 1.0,
				},
				ArgumentTwo: expr.ArgumentOne,
			}

		}

		return &Expression{
			Type: Operator,
			Value: "*",
			ArgumentOne: ExprInside,
			ArgumentTwo: ExprMain,
		}
	}

	return &Expression{
		Type: Variable,
		Value: "x",
		ArgumentOne: nil,
		ArgumentTwo: nil,
	}
}

func (expr *Expression) String() string {
	if expr.Type == Constant {
		return fmt.Sprint(expr.Value.(float64))
	} else if expr.Type == Variable {
		return expr.Value.(string)
	} else if expr.Type == Operator {
		return fmt.Sprintf("(%s %s %s)", expr.Value.(string), expr.ArgumentOne.String(), expr.ArgumentTwo.String())
	} else {
		return fmt.Sprintf("(%s %s)", expr.Value.(string), expr.ArgumentOne.String())
	}
}

func Diff(expression string) string {
	expr := GetExpression(expression)

	fmt.Printf("Before: %s | ", expr.String())

	if expr == nil {
		return ""
	}
	expr = expr.Diff()
	expr = expr.Reduce()

	fmt.Printf("After: %s\n", expr.String())
	return expr.String()
}

func GetExpression(expression string) *Expression {
	// complex
	if expression[0] == '(' {
		var exprType Type
		var value interface{}
		var argumentOne, argumentTwo string
		switch expression[1] {
		case '+','-','/','*','^':
			re := regexp.MustCompile(`^\(([+\-/*^])\s((?:-?[\d\w]+|\(.+?\)))\s((?:-?[\d\w]+|\(.+?\)))\)$`)
			matches := re.FindStringSubmatch(expression)
			exprType = Operator
			value = matches[1]
			argumentOne = matches[2]
			argumentTwo = matches[3]
		default:
			re := regexp.MustCompile(`^\((\w+)\s((?:-?[\d\w]+|\(.+?\)))\)$`)
			matches := re.FindStringSubmatch(expression)
			exprType = Function
			value = matches[1]
			argumentOne = matches[2]
		}

		if exprType == Operator {
			//fmt.Printf("%s is operator %s\n", expression, value)
			return &Expression{
				Type: Operator,
				Value: value,
				ArgumentOne: GetExpression(argumentOne),
				ArgumentTwo: GetExpression(argumentTwo),
			}
		}

		// fmt.Printf("%s is function\n", expression)
		return &Expression{
			Type: Function,
			Value: value,
			ArgumentOne: GetExpression(argumentOne),
			ArgumentTwo: nil,
		}
	} else {
		if expression == "x" {
			//fmt.Printf("%s is x\n", expression)
			return &Expression{
				Type: Variable,
				Value: "x",
				ArgumentOne: nil,
				ArgumentTwo: nil,
			}
		} else {
			//fmt.Printf("%s is const\n", expression)
			value, _ := strconv.ParseFloat(expression, 64)
			return &Expression{
				Type: Constant,
				Value: value,
				ArgumentOne: nil,
				ArgumentTwo: nil,
			}
		}
	}
}
