package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/schollz/progressbar/v3"
)

func testContracts(ctx *config.Context) {
	threads := 6
	offset := int64(0)

	for _, network := range ctx.Config.API.Networks {
		logger.Info("testing %s contract endpoints...", network)

		contracts, err := ctx.Contracts.GetMany(map[string]interface{}{
			"network": network,
		})
		if err != nil {
			logger.Errorf("testContracts: %s", err.Error())
			return
		}

		total := len(contracts)
		contracts = contracts[offset:]

		tasks := make(chan contract.Contract, len(contracts))
		for i := range contracts {
			tasks <- contracts[i]
		}

		counter := offset
		var wg sync.WaitGroup
		stop := make(chan struct{}, threads)

		for i := 0; i < threads; i++ {
			wg.Add(1)
			go testContract(tasks, stop, &counter, &wg)
		}

		wg.Add(1)
		go stopThread(threads, total, &counter, stop, &wg)

		wg.Wait()
	}
}

func stopThread(threads, total int, counter *int64, stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if counter != nil {
		bar := progressbar.NewOptions(total, progressbar.OptionSetPredictTime(false), progressbar.OptionClearOnFinish(), progressbar.OptionShowCount())

		for int64(total) > *counter {
			ctr := int(*counter)
			if err := bar.Set(ctr); err != nil {
				return
			}

			time.Sleep(100 * time.Millisecond)
		}

		if err := bar.Set(total); err != nil {
			return
		}
	}

	for i := 0; i < threads; i++ {
		stop <- struct{}{}
	}
}

func testContract(tasks chan contract.Contract, stop chan struct{}, counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			return
		case contract := <-tasks:
			prefix := fmt.Sprintf("contract/%s/%s", contract.Network, contract.Address)

			if err := request(prefix); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/code", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/operations", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/migrations", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/transfers", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/tokens", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/storage", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/same", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/similar", prefix)); err != nil {
				logger.Error(err)
			}
			if err := request(fmt.Sprintf("%s/entrypoints", prefix)); err != nil {
				logger.Error(err)
			}
			for i := range contract.Entrypoints {
				if err := request(fmt.Sprintf("%s/entrypoints/schema?entrypoint=%s", prefix, contract.Entrypoints[i])); err != nil {
					logger.Error(err)
				}
			}
			atomic.AddInt64(counter, 1)
		}

	}
}
