```mermaid
flowchart LR
    P[Producer] -->|jobs| CH[(Channel)]
    CH --> W1[Worker 1]
    CH --> W2[Worker 2]
    CH --> W3[Worker N]

    W1 -->|update| R[(Shared Result <br> protected by Mutex)]
    W2 -->|update| R
    W3 -->|update| R

