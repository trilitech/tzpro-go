package main

import (
    "context"
    "fmt"

    "blockwatch.cc/tzgo/tezos"
    "blockwatch.cc/tzpro-go"
)

func main() {
    if err := run(); err != nil {
        fmt.Println("Error:", err)
    }
}

func run() error {
    // use a placeholder calling context
    ctx := context.Background()

    // create a new SDK client
    c, err := tzpro.NewClient("https://api.staging.tzstats.com", nil)
    if err != nil {
        return err
    }

    // Params
    bp := tzpro.NewBlockParams().WithRights().WithMeta()
    op := tzpro.NewOpParams().WithStorage().WithMeta()
    ap := tzpro.NewAccountParams().WithMeta()
    bkp := tzpro.NewBakerParams().WithMeta()
    cp := tzpro.NewContractParams().WithMeta()

    // -----------------------------------------------------------------
    // Common
    //

    // fetch status
    stat, err := c.GetStatus(ctx)
    if err != nil {
        return err
    }
    if stat.Status != "synced" {
        return fmt.Errorf("Status is %s", stat.Status)
    }

    // tip
    tip, err := c.GetTip(ctx)
    if err != nil {
        return fmt.Errorf("Tip: %v", err)
    }

    // protocols
    if p, err := c.ListProtocols(ctx); err != nil || len(p) == 0 {
        return fmt.Errorf("ListProtocols: len=%d %v", len(p), err)
    }

    // config
    if _, err := c.GetConfig(ctx); err != nil {
        return fmt.Errorf("GetConfig: %v", err)
    }

    // config from height
    if _, err := c.GetConfigHeight(ctx, tip.Height); err != nil {
        return fmt.Errorf("GetConfigHeight: %v", err)
    }

    // -----------------------------------------------------------------
    // Block
    //

    // block
    if _, err := c.GetBlock(ctx, tip.Hash, bp); err != nil {
        return fmt.Errorf("GetBlock: %v", err)
    }

    // block head
    if _, err := c.GetHead(ctx, bp); err != nil {
        return fmt.Errorf("GetHead: %v", err)
    }

    // block height
    if _, err := c.GetBlockHeight(ctx, tip.Height, bp); err != nil {
        return fmt.Errorf("GetBlockHeight: %v", err)
    }

    // block with ops
    if b, err := c.GetBlockWithOps(ctx, tip.Hash, bp); err != nil || len(b.Ops) == 0 {
        return fmt.Errorf("GetBlockWithOps: len=%d %v", len(b.Ops), err)
    }

    // block ops
    if ops, err := c.GetBlockOps(ctx, tip.Hash, op); err != nil || len(ops) == 0 {
        return fmt.Errorf("GetBlockOps: len=%d %v", len(ops), err)
    }

    // block table
    bq := c.NewBlockQuery()
    bq.WithLimit(2).WithDesc()
    _, err = bq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Block query: %v", err)
    }

    // -----------------------------------------------------------------
    // Account
    //
    addr := tezos.MustParseAddress("tz1go7f6mEQfT2xX2LuHAqgnRGN6c2zHPf5c") // Main
    // addr := tezos.MustParseAddress("tz29WDVtnm7nQNS2GW45i3jPvKwa6Wkx7qos") // Ithaca

    // account
    if _, err := c.GetAccount(ctx, addr, ap); err != nil {
        return fmt.Errorf("GetAccount: %v", err)
    }

    // contracts
    if _, err := c.GetAccountContracts(ctx, addr, ap); err != nil {
        return fmt.Errorf("GetAccountContracts: %v", err)
    }

    // ops
    if ops, err := c.GetAccountOps(ctx, addr, op); err != nil || len(ops) == 0 {
        return fmt.Errorf("GetAccountOps: len=%d %v", len(ops), err)
    }

    // account table
    aq := c.NewBlockQuery()
    aq.WithLimit(2).WithDesc()
    _, err = aq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Account query: %v", err)
    }

    // -----------------------------------------------------------------
    // Baker
    //
    addr = tezos.MustParseAddress("tz1go7f6mEQfT2xX2LuHAqgnRGN6c2zHPf5c") // main
    // addr = tezos.MustParseAddress("tz1edUYGqBtteStneTGDBrQWTFmq9cnEELiW") // ithaca

    // baker
    if _, err := c.GetBaker(ctx, addr, bkp); err != nil {
        return fmt.Errorf("GetBaker: %v", err)
    }

    // list
    if l, err := c.ListBakers(ctx, bkp); err != nil || len(l) == 0 {
        return fmt.Errorf("ListBakers: len=%d %v", len(l), err)
    }

    // votes
    if ops, err := c.ListBakerVotes(ctx, addr, op); err != nil {
        return fmt.Errorf("ListBakerVotes: len=%d %v", len(ops), err)
    }

    // endorse
    if ops, err := c.ListBakerEndorsements(ctx, addr, op); err != nil {
        return fmt.Errorf("ListBakerEndorsements: len=%d %v", len(ops), err)
    }

    // deleg
    if ops, err := c.ListBakerDelegations(ctx, addr, op); err != nil {
        return fmt.Errorf("ListBakerDelegations: len=%d %v", len(ops), err)
    }

    // rights
    if _, err := c.ListBakerRights(ctx, addr, 400, bkp); err != nil {
        return fmt.Errorf("ListBakerRights: %v", err)
    }

    // income
    if _, err := c.GetBakerIncome(ctx, addr, 400, bkp); err != nil {
        return fmt.Errorf("GetBakerIncome: %v", err)
    }

    // snapshot
    if _, err := c.GetBakerSnapshot(ctx, addr, 400, bkp); err != nil {
        return fmt.Errorf("GetBakerSnapshot: %v", err)
    }

    // rights table
    rq := c.NewCycleRightsQuery()
    rq.WithLimit(2).WithDesc()
    _, err = rq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Rights query: %v", err)
    }

    // snapshot table
    sq := c.NewSnapshotQuery()
    sq.WithLimit(2).WithDesc()
    _, err = sq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Snapshot query: %v", err)
    }

    // -----------------------------------------------------------------
    // Bigmap
    //

    // allocs (find a bigmap with >0 keys)
    var bmid int64 = 1
    for {
        if bm, err := c.GetBigmap(ctx, bmid, cp); err != nil {
            return fmt.Errorf("GetBigmap: %v", err)
        } else if bm.NKeys > 0 {
            break
        }
        bmid++
    }

    // keys
    if k, err := c.ListBigmapKeys(ctx, bmid, cp); err != nil {
        return fmt.Errorf("ListBigmapKeys: %v", err)
    } else {
        if _, err := c.ListBigmapKeyUpdates(ctx, bmid, k[0].KeyHash.String(), cp); err != nil {
            return fmt.Errorf("ListBigmapKeyUpdates: %v", err)
        }
        // value
        if _, err := c.GetBigmapValue(ctx, bmid, k[0].KeyHash.String(), cp); err != nil {
            return fmt.Errorf("GetBigmapValue: %v", err)
        }
    }

    // list values
    if _, err := c.ListBigmapValues(ctx, bmid, cp); err != nil {
        return fmt.Errorf("ListBigmapValues: %v", err)
    }

    // list updates
    if _, err := c.ListBigmapUpdates(ctx, bmid, cp); err != nil {
        return fmt.Errorf("ListBigmapUpdates: %v", err)
    }

    // bigmap table
    bmq := c.NewBigmapQuery()
    bmq.WithLimit(2).WithDesc()
    _, err = bmq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Bigmap query: %v", err)
    }

    // bigmap update table
    bmuq := c.NewBigmapUpdateQuery()
    bmuq.WithLimit(2).WithDesc().WithFilter(tzpro.FilterModeEqual, "bigmap_id", bmid)
    _, err = bmuq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Bigmap update query: %v", err)
    }

    // bigmap value table
    bmvq := c.NewBigmapValueQuery()
    bmvq.WithLimit(2).WithDesc().WithFilter(tzpro.FilterModeEqual, "bigmap_id", bmid)
    _, err = bmvq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Bigmap value query: %v", err)
    }

    // -----------------------------------------------------------------
    // Chain
    //
    chq := c.NewChainQuery()
    chq.WithLimit(2).WithDesc()
    _, err = chq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Chain query: %v", err)
    }

    // -----------------------------------------------------------------
    // Constant
    //
    coq := c.NewConstantQuery()
    coq.WithLimit(2).WithDesc()
    _, err = coq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Constant query: %v", err)
    }

    // -----------------------------------------------------------------
    // Contract
    //
    addr = tezos.MustParseAddress("KT1EVPNZtekBirJhvALU5gNJS2F3ibWZXnpd") // main
    // addr = tezos.MustParseAddress("KT1UcwQtaztLSq8oufdXAtWpRTfFySCj7gFM") // ithaca

    // contract
    if _, err := c.GetContract(ctx, addr, cp); err != nil {
        return fmt.Errorf("GetContract: %v", err)
    }

    // script
    if _, err := c.GetContractScript(ctx, addr, cp); err != nil {
        return fmt.Errorf("GetContractScript: %v", err)
    }

    // storage
    if _, err := c.GetContractStorage(ctx, addr, cp); err != nil {
        return fmt.Errorf("GetContractStorage: %v", err)
    }

    // calls
    if _, err := c.ListContractCalls(ctx, addr, cp); err != nil {
        return fmt.Errorf("GetContractCalls: %v", err)
    }

    ccq := c.NewContractQuery()
    ccq.WithLimit(2).WithDesc()
    _, err = ccq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Contract query: %v", err)
    }

    // -----------------------------------------------------------------
    // Gov
    //
    electionId := 11 // main
    if _, err := c.GetElection(ctx, electionId); err != nil {
        return fmt.Errorf("GetElection: %v", err)
    }

    if _, err := c.ListVoters(ctx, electionId, 1); err != nil {
        return fmt.Errorf("ListVoters: %v", err)
    }

    if _, err := c.ListBallots(ctx, electionId, 1); err != nil {
        return fmt.Errorf("ListBallots: %v", err)
    }

    // -----------------------------------------------------------------
    // Operations
    //
    oq := c.NewOpQuery()
    oq.WithFilter(tzpro.FilterModeEqual, "type", "transaction").
        WithLimit(10).
        WithOrder(tzpro.OrderDesc)
    ores, err := oq.Run(ctx)
    if err != nil {
        return fmt.Errorf("Op query: %v", err)
    }
    if ores.Len() > 0 {
        if _, err := c.GetOp(ctx, ores.Rows[0].Hash, op); err != nil {
            return fmt.Errorf("GetOp: %v", err)
        }
    }

    fmt.Println("OK.")
    return nil
}
