package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"minecraft/cloudflare"
	"minecraft/docker"
)

type ServerRequest struct {
    Subdomain string `json:"subdomain"`
}

func CreateServer(w http.ResponseWriter, r *http.Request) {
    var req ServerRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ctx := context.Background()
    containerID, err := docker.CreateMinecraftServer(ctx, req.Subdomain)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Assuming the IP of your Railway service can be obtained somehow
    ip := "your-railway-service-ip"

    if err := cloudflare.CreateSubdomain(req.Subdomain, ip); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "containerID": containerID,
        "subdomain":   req.Subdomain + ".yourdomain.com",
    })
}
