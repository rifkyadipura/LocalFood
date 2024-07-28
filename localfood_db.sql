-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 28, 2024 at 06:39 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.0.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `localfood_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `artikel`
--

CREATE TABLE `artikel` (
  `id_artikel` int(11) NOT NULL,
  `id_users` int(11) NOT NULL,
  `judul` varchar(55) NOT NULL,
  `isi` varchar(255) NOT NULL,
  `foto` varchar(55) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `maps` varchar(255) NOT NULL,
  `status` enum('menunggu','selesai','ditolak') DEFAULT 'menunggu',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `artikel`
--

INSERT INTO `artikel` (`id_artikel`, `id_users`, `judul`, `isi`, `foto`, `alamat`, `maps`, `status`, `created_at`, `updated_at`) VALUES
(1, 2, 'Ayam Geprek 71', 'Restoran Masakan Ayam', '/uploads/Ayam Geprek 71.PNG', 'Sarijadi, Kec. Sukasari, Kota Bandung, Jawa Barat 40151', 'https://maps.app.goo.gl/pYz5dPkLEwGnpBXF6', 'selesai', '2024-07-27 21:31:56', '2024-07-28 04:35:04'),
(2, 3, 'Bakaran Jagonya Ayam Bakar', 'Restoran Indonesia', '/uploads/bakaran.PNG', 'Jl. Sariasih No.73, Sarijadi, Kec. Sukasari, Kota Bandung, Jawa Barat 40151', 'https://maps.app.goo.gl/iVYx1xVBXj5n2m8NA', 'menunggu', '2024-07-27 21:34:34', '2024-07-28 04:35:00');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_users` int(11) NOT NULL,
  `role` enum('admin','pengguna') DEFAULT 'pengguna',
  `email` varchar(155) NOT NULL,
  `username` varchar(155) NOT NULL,
  `password` varchar(155) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_users`, `role`, `email`, `username`, `password`, `created_at`, `updated_at`) VALUES
(1, 'admin', 'admin@gmail.com', 'admin', '41e5653fc7aeb894026d6bb7b2db7f65902b454945fa8fd65a6327047b5277fb', '2024-07-27 21:27:29', '2024-07-27 21:27:29'),
(2, 'pengguna', 'rifkyadipura@gmail.com', 'rifky', '3c8fe512459a840fbf6f4f69705b09cc0a35396c601a4de2d6b3098102fe816b', '2024-07-27 21:29:26', '2024-07-27 21:29:26'),
(3, 'pengguna', 'iqbaal@gmail.com', 'iqbaal', 'c92855f7c0cbbdc1abed5446d26c0052815621d40d963a0de97613922274c93c', '2024-07-27 21:33:15', '2024-07-27 21:33:15');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `artikel`
--
ALTER TABLE `artikel`
  ADD PRIMARY KEY (`id_artikel`),
  ADD KEY `fk_users` (`id_users`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_users`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `artikel`
--
ALTER TABLE `artikel`
  MODIFY `id_artikel` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_users` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `artikel`
--
ALTER TABLE `artikel`
  ADD CONSTRAINT `fk_users` FOREIGN KEY (`id_users`) REFERENCES `users` (`id_users`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
