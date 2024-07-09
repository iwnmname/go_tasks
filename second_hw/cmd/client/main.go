package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func main() {
	portVal := flag.Int("port", 1323, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newNameVal := flag.String("new_name", "", "new name for account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newNameVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		return create(cmd)
	case "get":
		return get(cmd)
	case "delete":
		return deleteAccount(cmd)
	case "patch":
		return patchAccount(cmd)
	case "change":
		return changeAccount(cmd)
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func create(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("Account %s created successfully with amount %d\n", cmd.Name, cmd.Amount)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func get(cmd Command) error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http get failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get account failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	var response dto.GetAccountResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("json unmarshal failed: %w", err)
	}

	fmt.Printf("Account name: %s, amount: %d\n", response.Name, response.Amount)

	return nil
}

func deleteAccount(cmd Command) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name), nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Printf("Account %s deleted successfully\n", cmd.Name)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func patchAccount(cmd Command) error {
	url := fmt.Sprintf("http://%s:%d/account?name=%s&amount=%d", cmd.Host, cmd.Port, cmd.Name, cmd.Amount)
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http patch failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	if resp.StatusCode == http.StatusOK {
		var account struct {
			Name   string `json:"name"`
			Amount int    `json:"amount"`
		}
		if err := json.Unmarshal(body, &account); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
		fmt.Printf("Account %s patched successfully. New amount: %d\n", account.Name, account.Amount)
	} else {
		return fmt.Errorf("patch account failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

func changeAccount(cmd Command) error {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s:%d/account?name=%s&new_name=%s", cmd.Host, cmd.Port, cmd.Name, cmd.NewName), nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http put failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Printf("Account name changed from %s to %s successfully\n", cmd.Name, cmd.NewName)

		getCmd := Command{
			Port: cmd.Port,
			Host: cmd.Host,
			Cmd:  "get",
			Name: cmd.NewName,
		}
		if err := get(getCmd); err != nil {
			return fmt.Errorf("failed to get updated account info: %w", err)
		}
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("change account failed: %s", string(body))
	}

	return nil
}
