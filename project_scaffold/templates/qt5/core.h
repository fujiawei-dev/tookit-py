{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__CORE_H
#define {{APP_NAME_UPPER}}__CORE_H

#include <QDebug>
#include <QObject>
#include <QSettings>
#include <QTimer>
#include <QWebSocket>

class Core : public QObject {
    Q_OBJECT

public:
    explicit Core(QObject *parent = nullptr);

    bool DebugMode() const;
    void InitConfig(bool, QSettings *);

    Q_PROPERTY(bool debugMode READ DebugMode);    // Read only
    Q_PROPERTY(QList<QString> items MEMBER items);// Read and write

    Q_INVOKABLE static QString getUuid();

    Q_INVOKABLE QString getRegionByCode(const QString &code);
    Q_INVOKABLE QList<QString> getProvinces();
    Q_INVOKABLE QList<QString> getCitiesByProvince(const QString &province);
    Q_INVOKABLE QList<QString> getDistrictsByProvinceCity(const QString &province, const QString &city);

    Q_INVOKABLE void connectToWebsocketServer(const QString &);
    Q_INVOKABLE void sendTextMessageToWebsocketServer(const QString &);

    static std::string AESEncryptStr(const QString &msgStr, const QString &keyStr);
    static std::string AESDecryptStr(const QString &msgStr, const QString &keyStr);


signals:
    void finished();// 正常退出
    void abort();   // 异常中断

public slots:
    void onExit();
    void onRun();// for console app

private Q_SLOTS:
    void onWebsocketTimeout();
    void onWebsocketConnected();
    void onWebsocketDisconnected();
    void onWebsocketTextMessageReceived(const QString &message);

private:
    bool debugMode = true;
    bool isExiting = false;// 表示即将退出程序

    void beforeInitConfig();
    void afterInitConfig();

    // variables from config file
    QSettings *conf{};
    QString remoteHostPort;
    QString remoteHttpBasePath;
    QString websocketPrefix;
    QList<QString> items = {};

    QWebSocket *websocketClient;
    QString websocketUrl;
    QTimer websocketTimer;

    void parseRegionDatabase();
    QMap<QString, QMap<QString, QList<QString>>> provinceCityDistrictMap;
    QMap<QString, QString> codeRegionMap;
};

#endif//{{APP_NAME_UPPER}}__CORE_H
