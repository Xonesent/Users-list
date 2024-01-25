package s_help

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"users-list/server"

	"github.com/spf13/viper"
)

func Enriche(ctx context.Context, name string) (*server.Enricher_structure, error) {
	err_ch := make(chan error)
	ret_ch := make(chan struct{})
	return_value := &server.Enricher_structure{}

	go func() {
		defer close(err_ch)
		defer close(ret_ch)
		
		wg := &sync.WaitGroup{}
		wg.Add(3)

		go func() {
			defer wg.Done()
			age, err := Get_age(ctx, name)
			if err != nil {
				err_ch <- err
				return
			}
			return_value.Age = age
		}()

		go func() {
			defer wg.Done()
			gender, err := Get_gender(ctx, name)
			if err != nil {
				err_ch <- err
				return
			}
			return_value.Gender = gender
		}()

		go func() {
			defer wg.Done()
			nationality, err := Get_nationality(ctx, name)
			if err != nil {
				err_ch <- err
				return
			}
			return_value.Nationality = nationality
		}()

		wg.Wait()
		ret_ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("server error")
	case err := <-err_ch:
		return nil, err
	case <-ret_ch:
		return return_value, nil
	}
}

func Get_nationality(ctx context.Context, name string) (string, error) {
	return_value := server.Nationalities_str{}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", viper.GetString("nation_api"), name), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("impossible to get age")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&return_value)
	if err != nil {
		return "", err
	}

	if len(return_value.Country) == 0 {
		return "", nil
	}
	return return_value.Country[0].CountryId, nil
}

func Get_gender(ctx context.Context, name string) (string, error) {
	return_value := &server.Gender_str{}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", viper.GetString("gender_api"), name), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("impossible to get age")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&return_value)
	if err != nil {
		return "", err
	}

	return return_value.Gender, nil
}

func Get_age(ctx context.Context, name string) (int, error) {
	return_value := &server.Age_str{}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", viper.GetString("age_api"), name), nil)
	if err != nil {
		return -1, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != 200 {
		return -1, fmt.Errorf("impossible to get age")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&return_value)
	if err != nil {
		return -1, err
	}

	return return_value.Age, nil
}
