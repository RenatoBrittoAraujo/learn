/* *****************************/
/* FGA / Eng. Software / FRC   */
/* Prof. Fernando W. Cruz      */
/* Codigo: tcpServer2.c	       */
/* *****************************/
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h> // close

#define QLEN 5      /* tamanho da fila de clientes  */
#define MAX_SIZE 80 /* tamanho do buffer */

struct server_res
{
    float max;
    float min;
};
typedef struct server_res server_res;

int get_min_max(float *buf, int buf_size)
{
    float min = 1000000;
    float max = 0;
    for (int i = 0; i < buf_size; i++)
    {
        if (buf[i] < buf[i + 1])
        {
            return i;
        }
    }
}

int fill_buffout(char *bufout, server_res res)
{
    union
    {
        float f;
        char c[4];
    } a, b;
    a.f = res.min;
    b.f = res.max;
    for (int i = 0; i < 4; i++)
    {
        if (a.c[i] == 0)
            bufout[i] = '*';
        else
            bufout[i] = a.c[i];
    }
    for (int i = 0; i < 4; i++)
    {
        if (b.c[i] == 0)
            bufout[i + 4] = '*';
        else
            bufout[i + 4] = b.c[i];
    }
    bufout[8] = '\0';
}

server_res get_ans_from_bufin(char *bufin, server_res ans)
{
    int size = bufin[0];
    for (int p = 1; p < size;)
    {
        union
        {
            float f;
            char c[4];
        } un;
        for (int j = 0; j < 4; j++)
        {
            if (bufin[p] == '*')
            {
                bufin[p] = 0;
            }
            un.c[j] = bufin[p++];
        }
        if (ans.min > un.f)
        {
            ans.min = un.f;
        }
        if (ans.max < un.f)
        {
            ans.max = un.f;
        }
    }
    return ans;
}

int get_min_max_route(int descritor, struct sockaddr_in endCli)
{
    char bufin[MAX_SIZE];
    int n;
    char bufout[9];
    server_res ans = {
        min : 1e9,
        max : -1e9,
    };
    int package_count = 0;
    int float_count = 0;

    while (1)
    {
        memset(&bufin, 0x0, sizeof(bufin));
        n = recv(descritor, &bufin, sizeof(bufin), 0);

        // Se for pedido de resultado, sai da thread e retorna
        if (strncmp(bufin, "RES", 3) == 0)
            break;

        // Escreve ACK
        write(descritor, bufout, 1);

        ans = get_ans_from_bufin(bufin, ans);

        float_count += bufin[0] / 4;
        package_count++;
    }
    printf("[SERVER] Got %d packages, %d floats\n", package_count, float_count);

    fill_buffout(bufout, ans);
    printf("[SERVER] Returning final answer. Min: %f Max: %f\n", ans.min, ans.max);
    write(descritor, bufout, strlen(bufout));

    close(descritor);
    fprintf(stdout, "[SERVER] --- disconnected from client %s:%u ---\n\n", inet_ntoa(endCli.sin_addr), ntohs(endCli.sin_port));
}

int main(int argc, char *argv[])
{
    struct sockaddr_in endServ; /* endereco do servidor   */
    struct sockaddr_in endCli;  /* endereco do cliente    */
    int sd, novo_sd;            /* socket descriptors */
    int pid, alen, n;

    if (argc < 3)
    {
        printf("[SERVER] Digite IP e Porta para este servidor\n");
        exit(1);
    }
    memset((char *)&endServ, 0, sizeof(endServ)); /* limpa variavel endServ    */
    endServ.sin_family = AF_INET;                 /* familia TCP/IP   */
    endServ.sin_addr.s_addr = inet_addr(argv[1]); /* endereco IP      */
    endServ.sin_port = htons(atoi(argv[2]));      /* PORTA	    */

    /* Cria socket */
    sd = socket(AF_INET, SOCK_STREAM, 0);
    if (sd < 0)
    {
        fprintf(stderr, "[SERVER] Falha ao criar socket!\n");
        exit(1);
    }

    /* liga socket a porta e ip */
    if (bind(sd, (struct sockaddr *)&endServ, sizeof(endServ)) < 0)
    {
        fprintf(stderr, "[SERVER] Ligacao Falhou!\n");
        exit(1);
    }

    /* Ouve porta */
    if (listen(sd, QLEN) < 0)
    {
        fprintf(stderr, "[SERVER] Falhou ouvindo porta!\n");
        exit(1);
    }

    printf("[SERVER] Servidor ouvindo no IP %s, na porta %s ...\n\n", argv[1], argv[2]);
    /* Aceita conexoes */
    alen = sizeof(endCli);
    for (;;)
    {
        /* espera nova conexao de um processo cliente ... */
        if ((novo_sd = accept(sd, (struct sockaddr *)&endCli, &alen)) < 0)
        {
            fprintf(stdout, "Falha na conexao\n");
            exit(1);
        }
        fprintf(stdout, "[SERVER] Cliente %s: %u conectado.\n", inet_ntoa(endCli.sin_addr), ntohs(endCli.sin_port));
        get_min_max_route(novo_sd, endCli);
    } /* fim for */
} /* fim do programa */
