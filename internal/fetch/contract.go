package fetch

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/baking-bad/bcdhub/internal/bcd"
	"github.com/pkg/errors"
)

const (
	contractFormatPath = "%s/contracts/%s/%s_%s.json"
)

var (
	delegatorContract = []byte(`[{"prim":"parameter","args":[{"prim":"or","args":[{"prim":"lambda","args":[{"prim":"unit"},{"prim":"list","args":[{"prim":"operation"}]}],"annots":["%do"]},{"prim":"unit","annots":["%default"]}]}]},{"prim":"storage","args":[{"prim":"key_hash"}]},{"prim":"code","args":[[[[{"prim":"DUP"},{"prim":"CAR"},{"prim":"DIP","args":[[{"prim":"CDR"}]]}]],{"prim":"IF_LEFT","args":[[{"prim":"PUSH","args":[{"prim":"mutez"},{"int":"0"}]},{"prim":"AMOUNT"},[[{"prim":"COMPARE"},{"prim":"EQ"}],{"prim":"IF","args":[[],[[{"prim":"UNIT"},{"prim":"FAILWITH"}]]]}],[{"prim":"DIP","args":[[{"prim":"DUP"}]]},{"prim":"SWAP"}],{"prim":"IMPLICIT_ACCOUNT"},{"prim":"ADDRESS"},{"prim":"SENDER"},[[{"prim":"COMPARE"},{"prim":"EQ"}],{"prim":"IF","args":[[],[[{"prim":"UNIT"},{"prim":"FAILWITH"}]]]}],{"prim":"UNIT"},{"prim":"EXEC"},{"prim":"PAIR"}],[{"prim":"DROP"},{"prim":"NIL","args":[{"prim":"operation"}]},{"prim":"PAIR"}]]}]]}]`)
)

// RemoveContract -
func RemoveContract(address, network, protocol, filesDirectory string) error {
	if filesDirectory == "" {
		return errors.Errorf("Invalid filesDirectory: %s", filesDirectory)
	}
	protoSymLink, err := bcd.GetProtoSymLink(protocol)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf(contractFormatPath, filesDirectory, network, address, protoSymLink)
	if _, err = os.Stat(filePath); err == nil {
		return os.Remove(filePath)
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}

// RemoveAllContracts -
func RemoveAllContracts(network, filesDirectory string) error {
	if filesDirectory == "" {
		return errors.Errorf("Invalid filesDirectory: %s", filesDirectory)
	}

	dirPath := fmt.Sprintf("%s/contracts/%s", filesDirectory, network)
	if _, err := os.Stat(dirPath); err == nil {
		return os.RemoveAll(dirPath)
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Contract - reads contract from file system
func Contract(address, network, protocol, filesDirectory string) ([]byte, error) {
	if protocol == "" {
		protocol = bcd.GetCurrentProtocol()
	}
	protoSymLink, err := bcd.GetProtoSymLink(protocol)
	if err != nil {
		return nil, err
	}

	filePath := fmt.Sprintf(contractFormatPath, filesDirectory, network, address, protoSymLink)
	if _, err = os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return delegatorContract, nil
		} else {
			return nil, err
		}
	}
	return ioutil.ReadFile(filePath)
}
