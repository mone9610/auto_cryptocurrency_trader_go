-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- ホスト: mysql
-- 生成日時: 2021 年 1 月 24 日 05:57
-- サーバのバージョン： 5.7.32
-- PHP のバージョン: 7.4.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- データベース: `auto_trader`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `buy_order`
--

CREATE TABLE `buy_order` (
  `id` int(11) NOT NULL,
  `order_type` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` float NOT NULL,
  `size` float NOT NULL,
  `child_order_id` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `execution_status` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- テーブルの構造 `conf`
--

CREATE TABLE `conf` (
  `id` int(11) NOT NULL,
  `access_key` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `secret_key` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_ready` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- テーブルのデータのダンプ `conf`
--

INSERT INTO `conf` (`id`, `access_key`, `secret_key`, `is_ready`) VALUES
(1, '', '', 0);

-- --------------------------------------------------------

--
-- テーブルの構造 `mode`
--

CREATE TABLE `mode` (
  `id` int(11) NOT NULL,
  `trade_mode` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- テーブルのデータのダンプ `mode`
--

INSERT INTO `mode` (`id`, `trade_mode`, `created_at`, `updated_at`) VALUES
(1, 0, '2020-12-29 09:39:16', '2021-01-17 08:24:03');

-- --------------------------------------------------------

--
-- テーブルの構造 `sell_order`
--

CREATE TABLE `sell_order` (
  `id` int(11) NOT NULL,
  `order_type` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` float NOT NULL,
  `size` float NOT NULL,
  `child_order_id` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `execution_status` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- テーブルの構造 `ticker_info`
--

CREATE TABLE `ticker_info` (
  `id` int(11) NOT NULL,
  `high` float NOT NULL,
  `last` float NOT NULL,
  `low` float NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- ダンプしたテーブルのインデックス
--

--
-- テーブルのインデックス `buy_order`
--
ALTER TABLE `buy_order`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `conf`
--
ALTER TABLE `conf`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `mode`
--
ALTER TABLE `mode`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `sell_order`
--
ALTER TABLE `sell_order`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `ticker_info`
--
ALTER TABLE `ticker_info`
  ADD PRIMARY KEY (`id`);

--
-- ダンプしたテーブルの AUTO_INCREMENT
--

--
-- テーブルの AUTO_INCREMENT `buy_order`
--
ALTER TABLE `buy_order`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- テーブルの AUTO_INCREMENT `conf`
--
ALTER TABLE `conf`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- テーブルの AUTO_INCREMENT `mode`
--
ALTER TABLE `mode`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- テーブルの AUTO_INCREMENT `sell_order`
--
ALTER TABLE `sell_order`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- テーブルの AUTO_INCREMENT `ticker_info`
--
ALTER TABLE `ticker_info`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
