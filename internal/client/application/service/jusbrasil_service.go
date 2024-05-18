package applicationsvc

import (
	"fmt"
	"github.com/augustomcosta/client-api/internal/client/domain/entity"
)

func VerifyJusbrasilHistory(client *entity.Client) error {
	if client.Name == "Amálio Benjamin Espíndola" {
		return fmt.Errorf("Você possui um histórico de processos muito extenso. Não foi possível")
	}
	return nil
}
