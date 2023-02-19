# Objective
- Peer-To-Peer storage

<details>
<summary>001</summary>

# Initialize New TCP Connection with port :3000
- You need Address for listening the peers
- You need 3 way handshake for TCP Connection (Not finished)

# Structure
- ListenAndAccept(Listener)
- - startAcceptLoop (Concurrency)
- - - handleConn (Connect, handshake)

# QnA
- Goroutine for handshakeFunc inside handleConn (Concurrency?)

</details>

# CMD
- <code>make run</code>
- <code>make test</code>

# Resources
- [What is the TCP 3-Way Handshake and Why Backend Engineers should understand it](https://www.youtube.com/watch?v=bW_BILl7n0Y)