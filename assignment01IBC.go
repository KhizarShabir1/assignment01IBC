package assignment01IBC
import
(

  "crypto/sha256"
  "encoding/hex"
  "fmt"
  // "reflect"
)

type Block struct {
transaction string
prevPointer *Block
hashVal string
}

func InsertBlock(transaction string, chainHead *Block) *Block {

  if chainHead==nil{
    block := Block{transaction, nil,"emp"}
    chainHead:= & block
    return chainHead
  } else{

    newBlock := Block{transaction, nil,"emp"}
    prev :=chainHead
    curr:=chainHead
    // xt:=reflect.TypeOf(curr).Kind()
    // fmt.Printf("%T: %s\n", xt, xt)

      for curr!=nil{
        prev=curr
        curr=curr.prevPointer
    	}

    prev.prevPointer=&newBlock
    hash := sha256.New()
    hash.Write([]byte(prev.transaction))
    prev.prevPointer.hashVal= hex.EncodeToString(hash.Sum(nil))
    return chainHead
  }
}
func ListBlocks(chainHead *Block) {

  curr:=chainHead
  fmt.Printf("\n")
  for curr!=nil{

    fmt.Printf(curr.transaction)
    fmt.Printf("  -> ")

    curr=curr.prevPointer
  }
  fmt.Printf("\n")

}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
  curr:=chainHead
  for curr!=nil{

    if curr.transaction==oldTrans{
      curr.transaction=newTrans
    }

    curr=curr.prevPointer
  }

}
func VerifyChain(chainHead *Block) {

  curr:=chainHead

  check:= true

  for curr!=nil{

    hash := sha256.New()
    hash.Write([]byte(curr.transaction))
    hashStr:= hex.EncodeToString(hash.Sum(nil))
    if curr.prevPointer!=nil{
      if hashStr!=curr.prevPointer.hashVal{
        check= false
      }
    }


    curr=curr.prevPointer
  }
  if check{

    fmt.Printf("\nBlockChain verified !!!\n")
  }else{
    fmt.Printf("\nBlockChain NOT verified !!!\n")
  }


}
